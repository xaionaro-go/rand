package main

import (
	"strings"
	"text/template"
)

var (
	templateRead = `
// Read{{ .MethodName }} is an analog of math/rand.Read. This random numbers could easily
// be predicted (it's not an analog of crypto/rand.Read).
//
// Applied PRNG method:
// {{ .AdditionalInfo }}
func (prng *PRNG) Read{{ .MethodName }}(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{{- if eq .ResultSize 8 }}
	{
		{{ .InitCode }}
		for i = 8; i < l; i += 8 {
			{{ .GetValueCode }}
			*(*uint64)((unsafe.Pointer)(&b[i-8])) = {{ .ResultVariable }}
		}
		{{ .FinishCode }}
	}
	i -= 8
	rest := prng.{{ .MethodName }}()
	if i == l {
		return
	}
	if l-i >= 4 {
		*(*uint32)((unsafe.Pointer)(&b[i])) = uint32(rest)
		rest >>= 32
		i += 4
	}
	{{- else }}
	{
		{{ .InitCode }}
		for i = 4; i < l; i += 4 {
			{{ .GetValueCode }}
			*(*uint32)((unsafe.Pointer)(&b[i-4])) = {{ .ResultVariable }}
		}
		{{ .FinishCode }}
	}
	i -= 4
	rest := prng.{{ .MethodName }}()
	if i == l {
		return
	}
	{{- end }}
	if l-i >= 2 {
		*(*uint16)((unsafe.Pointer)(&b[i])) = uint16(rest)
		rest >>= 16
		i += 2
	}
	if l-i >= 1 {
		b[i] = uint8(rest)
	}
	return
}

// XORRead XORs argument "b" with a pseudo-random value. The result is
// the same (but faster) as:
//
// 	x := make([]byte, len(b))
// 	mathrand.Read{{ .MethodName }}(x)
// 	for i := range b {
// 		b[i] ^= x[i]
// 	}
func (prng *PRNG) XORRead{{ .MethodName }}(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{{- if eq .ResultSize 8 }}
	{
		{{ .InitCode }}
		for i = 8; i < l; i += 8 {
			{{ .GetValueCode }}
			*(*uint64)((unsafe.Pointer)(&b[i-8])) ^= {{ .ResultVariable }}
		}
		{{ .FinishCode }}
	}
	i -= 8
	rest := prng.{{ .MethodName }}()
	if i == l {
		return
	}
	if l-i >= 4 {
		*(*uint32)((unsafe.Pointer)(&b[i])) = uint32(rest)
		rest >>= 32
		i += 4
	}
	{{- else }}
	{
		{{ .InitCode }}
		for i = 4; i < l; i += 4 {
			{{ .GetValueCode }}
			*(*uint32)((unsafe.Pointer)(&b[i-4])) = {{ .ResultVariable }}
		}
		{{ .FinishCode }}
	}
	i -= 4
	rest := prng.{{ .MethodName }}()
	if i == l {
		return
	}
	{{- end }}
	if l-i >= 2 {
		*(*uint16)((unsafe.Pointer)(&b[i])) = uint16(rest)
		rest >>= 16
		i += 2
	}
	if l-i >= 1 {
		b[i] = uint8(rest)
	}
	return
}
`
	templateTestRead = `
func TestRead{{ .MethodName }}(t *testing.T) {
	//testRead(t, mathrand.GlobalPRNG.Read{{ .MethodName }})
}
func BenchmarkRead{{ .MethodName }}1(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.Read{{ .MethodName }}, 1)
}
func BenchmarkRead{{ .MethodName }}16(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.Read{{ .MethodName }}, 16)
}
func BenchmarkRead{{ .MethodName }}1024(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.Read{{ .MethodName }}, 1024)
}
func BenchmarkRead{{ .MethodName }}65536(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.Read{{ .MethodName }}, 65536)
}
func BenchmarkRead{{ .MethodName }}65536Concurrent(b *testing.B) {
	benchmarkConcurrentRead(b, mathrand.GlobalPRNG.Read{{ .MethodName }}, 65536)
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
