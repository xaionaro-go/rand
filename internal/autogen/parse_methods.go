package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"strings"
)

type Method struct {
	Name         string
	InitCode     string
	GetValueCode string
	ResultVariable string
	FinishCode   string
	ResultSize   uint

	genInfo *genInfo
	initCodeLines []string
	getValueCodeLines []string
	finishCodeLines []string
	variableExists map[string]struct{}
}
type Methods []*Method

func ParseMethods() (methods Methods, err error) {
	for _, filePath := range []string{`uint64.go`, `uint32.go`} {
		var newMethods Methods
		newMethods, err = parseMethodsFromFile(filePath)
		if err != nil {
			return
		}
		methods = append(methods, newMethods...)
	}

	return
}

func parseMethodsFromFile(path string) (methods Methods, err error) {
	fileSet := token.NewFileSet()
	parsedFile, err := parser.ParseFile(fileSet, path, nil, 0)
	if err != nil {
		return
	}
	var genInfo *genInfo
	for _, decl := range parsedFile.Decls {
		switch decl := decl.(type) {
		case *ast.GenDecl:
			genInfo = parseGenInfo(decl)
		case *ast.FuncDecl:
			method := parseMethod(decl, genInfo)
			if method == nil {
				continue
			}
			methods = append(methods, method)
		default:
			panic(fmt.Sprintf("%T:%v", decl, decl))
		}
	}

	return
}

type genInfo struct {
	constMap map[string]ast.Expr
	funcMap map[string]*ast.FuncDecl
}

func parseGenInfo(genDecl *ast.GenDecl) (result *genInfo) {
	result = &genInfo{
		constMap: map[string]ast.Expr{},
		funcMap: map[string]*ast.FuncDecl{},
	}
	for _, spec := range genDecl.Specs {
		switch spec := spec.(type) {
		case *ast.ValueSpec:
			if len(spec.Names) != 1 {
				panic(spec.Names)
			}
			if len(spec.Values) != 1 {
				panic(spec.Values)
			}
			result.constMap[spec.Names[0].String()] = spec.Values[0]
		default:
			panic(fmt.Sprintf("%T:%v\n", spec, spec))
		}
	}

	return
}

func parseMethod(funcDecl *ast.FuncDecl, genInfo *genInfo) (method *Method) {
	if funcDecl.Recv == nil { // it's not a method, just a function (in the global scope)
		genInfo.funcMap[funcDecl.Name.String()] = funcDecl
		return
	}

	recvPtrType, ok := funcDecl.Recv.List[0].Type.(*ast.StarExpr)
	if !ok {
		return
	}
	recvType, ok := recvPtrType.X.(*ast.Ident)
	if !ok {
		return
	}
	if recvType.String() != `PRNG` {
		return
	}

	method = &Method{
		Name: funcDecl.Name.String(),
		genInfo: genInfo,
		variableExists: map[string]struct{}{},
	}

	resultTypeName := funcDecl.Type.Results.List[0].Type.(*ast.Ident).String()
	switch resultTypeName {
	case `uint32`:
		method.ResultSize = 4
	case `uint64`:
		method.ResultSize = 8
	}

	if len(funcDecl.Type.Results.List[0].Names) == 1 {
		method.ResultVariable = funcDecl.Type.Results.List[0].Names[0].String()
		method.initCodeLines = append(method.initCodeLines, `var `+method.ResultVariable+` `+resultTypeName)
	}

	var codeLines []string

	for _, decl := range funcDecl.Body.List {
		switch decl := decl.(type) {
		case *ast.AssignStmt:
			method.addAssignStmt(decl)
		case *ast.ReturnStmt:
			method.addReturnStmt(decl)
		default:
			panic(fmt.Sprintf("%T", decl))
		}
	}

	method.GetValueCode = strings.Join(codeLines, "\n")

	method.compile()
	return
}

func (m *Method) compile() {
	m.InitCode = strings.Join(m.initCodeLines, "\n\t\t")
	m.GetValueCode = strings.Join(m.getValueCodeLines, "\n\t\t\t")
	m.FinishCode = strings.Join(m.finishCodeLines, "\n\t\t")
}

func (m *Method) addReturnStmt(stmt *ast.ReturnStmt) {
	if len(stmt.Results) == 0 {
		return
	}
	if len(stmt.Results) != 1 {
		panic(stmt.Results)
	}
	m.ResultVariable = m.getCodeString(stmt.Results[0])
}

func (m *Method) addAssignStmt(stmt *ast.AssignStmt) {
	if len(stmt.Lhs) != 1 {
		panic(len(stmt.Lhs))
	}
	if len(stmt.Rhs) != 1 {
		panic(len(stmt.Rhs))
	}
	l := m.getCodeString(stmt.Lhs[0])
	r := m.getCodeString(stmt.Rhs[0])
	m.getValueCodeLines = append(m.getValueCodeLines, l+` `+stmt.Tok.String()+` `+r)
}

func (m *Method) getCodeString(expr ast.Expr) string {
	switch expr := expr.(type) {
	case *ast.Ident:
		v := expr.String()
		if expr.Obj != nil && expr.Obj.Kind.String() == `var` {
			m.considerVariable(v)
		}
		return v
	case *ast.SelectorExpr:
		if expr.X.(*ast.Ident).String() != `prng` {
			panic(expr.X)
		}
		v := expr.Sel.String() + "Temp"
		m.considerVariable(v)
		return v
	case *ast.IndexExpr:
		variable := m.getCodeString(expr.X) + strings.ToTitle(m.getCodeString(expr.Index))
		m.considerVariable(variable)
		return variable
	case *ast.CallExpr:
		var args []string
		for _, arg := range expr.Args {
			args = append(args, m.getCodeString(arg))
		}
		return m.getCodeString(expr.Fun) + `(` + strings.Join(args, `, `) + `)`
	case *ast.BasicLit:
		return expr.Value
	case *ast.BinaryExpr:
		y := m.getCodeString(expr.Y)
		return m.getCodeString(expr.X) + ` ` + expr.Op.String() + ` ` + y
	default:
		panic(fmt.Sprintf("%T", expr))
	}
}

func (m *Method) considerVariable(variable string) {
	if _, ok := m.variableExists[variable]; ok {
		return
	}
	m.variableExists[variable] = struct{}{}
	/*if m.genInfo.constMap[variable] != nil {
		return
	}*/
	if m.genInfo.funcMap[variable] != nil {
		return
	}
	if strings.HasPrefix(variable, `state64Temp`) {
		idxString := variable[len(`state64Temp`):]
		if len(idxString) == 0 {
			return
		}
		idx, err := strconv.ParseUint(idxString, 10, 64)
		if err != nil {
			panic(variable)
		}
		m.initCodeLines = append(m.initCodeLines, fmt.Sprintf(`%v := prng.state64[%v]`, variable, idx))
		m.finishCodeLines = append(m.finishCodeLines, fmt.Sprintf(`prng.state64[%v] = %v`, idx, variable))
		return
	}
	if strings.HasPrefix(variable, `state32Temp`) {
		idxString := variable[len(`state32Temp`):]
		if len(idxString) == 0 {
			return
		}
		idx, err := strconv.ParseUint(idxString, 10, 64)
		if err != nil {
			panic(err)
		}
		m.initCodeLines = append(m.initCodeLines, fmt.Sprintf(`%v := prng.state32[%v]`, variable, idx))
		m.finishCodeLines = append(m.finishCodeLines, fmt.Sprintf(`prng.state32[%v] = %v`, idx, variable))
		return
	}
	if variable == `pcgStateTemp` {
		m.initCodeLines = append(m.initCodeLines, fmt.Sprintf(`%v := prng.pcgState`, variable))
		m.finishCodeLines = append(m.finishCodeLines, fmt.Sprintf(`prng.pcgState = %v`, variable))
	}
}