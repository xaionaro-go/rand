package main

import (
	"strings"
	"text/template"
)

var (
	templateRead = `
{{- if .IsXORRead }}
// XORRead{{ .MethodName }}{{- if .EnableReseed }}WithReseed{{- end }} XORs argument "b" with a pseudo-random value.
// The result is the same (but faster) as:
//
// 	x := make([]byte, len(b))
// 	mathrand.Read{{ .MethodName }}{{- if .EnableReseed }}WithReseed{{- end }}(x)
// 	for i := range b {
// 		b[i] ^= x[i]
// 	}
{{- else }}
// Read{{ .MethodName }}{{- if .EnableReseed }}WithReseed{{- end }} is an analog of math/rand.Read. This random
// numbers could easily be predicted (it's not an analog of crypto/rand.Read).
{{- if .EnableReseed }}
//
// "Reseed" forces to use a new good seed on setting value to a pointer `+"`& 0xff == 0`"+`. It allows
// to improve randomness of random numbers with a small performance impact.
// This method makes sense only if len(b) is large enough (>= 256 bytes). Otherwise it could affect
// strongly performance or it will not improve the randomness.
{{- end }}
//
// Applied PRNG method:
// {{ .AdditionalInfo }}
{{- end }}
func (prng *PRNG) {{ if .IsXORRead }}XOR{{ end }}Read{{ .MethodName }}{{- if .EnableReseed }}WithReseed{{- end }}(b []byte) (l int, err error) {
	l = len(b)
	s := (uintptr)((unsafe.Pointer)(&b[0]))
	p := s
	e := s + uintptr(l)
	{
`+		// uint32 PRNG:
		//
		// len(b) == 12:  ....p...f...e   last p -> f
		//                ^   ^   ^   ^
		//
		// len(b) == 11:  ....p..f...e    last p -> f-3
		//                ^   ^   ^
		//
		// len(b) == 10:  ....p.f...e     last p -> f-2
		//                ^   ^   ^
		//
		// len(b) == 9:   ....pf...e      last p -> f-1
		//                ^   ^   ^
`		{{ .InitCode }}
		for f := e - {{ .ResultSize }}; p <= f; p += {{ .ResultSize }} {
			{{ .GetValueCode }}
			*(*uint{{ umul .ResultSize 8 }})((unsafe.Pointer)(p)) {{ if .IsXORRead }}^{{ end }}= {{ .ResultVariable }}
			{{- if .EnableReseed }}
			if p & 0xff < {{ .ResultSize }} {
				continue
			}
			{{- if eq .MethodName "Uint32PCG" }}
			pcgStateTemp = xorShift64(pcgStateTemp ^ primeNumber64bit0) << 1 + 1
			{{- else }}
			{{- if eq .ResultSize 8 }}
			state64Temp0 = xorShift64(state64Temp0 ^ primeNumber64bit0)
			{{- else }}
			state{{ umul .ResultSize 8 }}Temp0 = (uint{{ umul .ResultSize 8 }})(xorShift32(state{{ umul .ResultSize 8 }}Temp0 ^ primeNumber{{ umul .ResultSize 8 }}bit0))
			{{- end }}
			{{- end }}
			{{- end }}
		}
		{{ .FinishCode }}
	}

	if e - p == 0 {
		return
	}
	rest := prng.{{ .MethodName }}()

	{{- if eq .ResultSize 8 }}
	if e - p >= 4 {
		*(*uint32)((unsafe.Pointer)(p)) {{ if .IsXORRead }}^{{ end }}= uint32(rest)
		rest >>= 32
		p += 4
	}
	{{- end }}
	if e - p >= 2 {
		*(*uint16)((unsafe.Pointer)(p)) {{ if .IsXORRead }}^{{ end }}= uint16(rest)
		rest >>= 16
		p += 2
	}
	if e - p >= 1 {
		*(*uint8)((unsafe.Pointer)(p)) {{ if .IsXORRead }}^{{ end }}= uint8(rest)
	}
	return
}
`
	templateTestRead = `
func Test{{ if .IsXORRead }}XOR{{ end }}Read{{ .MethodName }}{{- if .EnableReseed }}WithReseed{{- end }}(t *testing.T) {
	//testRead(t, mathrand.GlobalPRNG.{{ if .IsXORRead }}XOR{{ end }}Read{{ .MethodName }}{{- if .EnableReseed }}WithReseed{{- end }})
	prng := mathrand.NewWithSeed(initialSeed)
	prepareSample(
		"{{ if .IsXORRead }}XOR{{ end }}{{ .MethodName }}{{- if .EnableReseed }}WithReseed{{- end }}",
		prng.Read{{ .MethodName }}{{- if .EnableReseed }}WithReseed{{- end }},
		{{ .EnableReseed }},
	)
    {{- if .EnableReseed }}
	a := make([]byte, 65536)
	b := make([]byte, 65536)
	mathrand.GlobalPRNG.{{ if .IsXORRead }}XOR{{ end }}Read{{ .MethodName }}(a)
	mathrand.GlobalPRNG.{{ if .IsXORRead }}XOR{{ end }}Read{{ .MethodName }}(b)
	assert.NotEqual(t, a, b)
    {{- end }}
}
func Benchmark{{ if .IsXORRead }}XOR{{ end }}Read{{ .MethodName }}{{- if .EnableReseed }}WithReseed{{- end }}1(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.{{ if .IsXORRead }}XOR{{ end }}Read{{ .MethodName }}{{- if .EnableReseed }}WithReseed{{- end }}, 1)
}
func Benchmark{{ if .IsXORRead }}XOR{{ end }}Read{{ .MethodName }}{{- if .EnableReseed }}WithReseed{{- end }}16(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.{{ if .IsXORRead }}XOR{{ end }}Read{{ .MethodName }}{{- if .EnableReseed }}WithReseed{{- end }}, 16)
}
func Benchmark{{ if .IsXORRead }}XOR{{ end }}Read{{ .MethodName }}{{- if .EnableReseed }}WithReseed{{- end }}1024(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.{{ if .IsXORRead }}XOR{{ end }}Read{{ .MethodName }}{{- if .EnableReseed }}WithReseed{{- end }}, 1024)
}
func Benchmark{{ if .IsXORRead }}XOR{{ end }}Read{{ .MethodName }}{{- if .EnableReseed }}WithReseed{{- end }}65536(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.{{ if .IsXORRead }}XOR{{ end }}Read{{ .MethodName }}{{- if .EnableReseed }}WithReseed{{- end }}, 65536)
}
func Benchmark{{ if .IsXORRead }}XOR{{ end }}Read{{ .MethodName }}{{- if .EnableReseed }}WithReseed{{- end }}65536Concurrent(b *testing.B) {
	benchmarkConcurrentRead(b, mathrand.GlobalPRNG.{{ if .IsXORRead }}XOR{{ end }}Read{{ .MethodName }}{{- if .EnableReseed }}WithReseed{{- end }}, 65536)
}
`
	templateTestMethod = `
func Test{{ .MethodName }}(t *testing.T) {
	{{- if eq .ResultSize 8 }}
	//testUint64(t, mathrand.GlobalPRNG.{{ .MethodName }})
	{{- else }}
	//testUint32(t, mathrand.GlobalPRNG.{{ .MethodName }})
	{{- end }}
}
func Benchmark{{ .MethodName }}(b *testing.B) {
	prng := mathrand.New()
	b.SetBytes({{ .ResultSize }})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prng.{{ .MethodName }}()
	}
}
`
)

type Templates struct {
	TestMethod *template.Template
	Read       *template.Template
	TestRead   *template.Template
}

func GetTemplates() (result *Templates, err error) {
	result = &Templates{}

	funcs := map[string]interface{}{
		"stringsJoin": func(slice []string, sep string) string {
			return strings.Join(slice, sep)
		},
		"umul": func(a, b uint) uint {
			return a * b
		},
	}

	result.TestMethod, err = template.New(`TestMethod`).
		Funcs(funcs).Parse(templateTestMethod)
	if err != nil {
		return
	}

	result.Read, err = template.New(`Read`).
		Funcs(funcs).Parse(templateRead)
	if err != nil {
		return
	}

	result.TestRead, err = template.New(`TestRead`).
		Funcs(funcs).Parse(templateTestRead)
	if err != nil {
		return
	}

	return
}
