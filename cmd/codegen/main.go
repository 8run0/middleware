package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type IInfo struct {
	PackageName string
	IName       string
	CName       string
	IFuncs      []IFunc
}

type IFunc struct {
	IName     string
	FName     string
	InParams  []FParam
	OutParams []FParam
}

type FParam struct {
	PName string
	PType string
}

func main() {
	interfaceStr :=
		`package {{ .PackageName }}

	type {{.IName}}Impl interface {
		{{range .IFuncs}}{{.FName}} ({{range .InParams}} {{.PName}} {{.PType}},{{end}}) ({{range .OutParams}}{{.PName}} {{.PType}},{{end}})
		{{end}}
	}
	`

	starterStr :=
		`package {{ .PackageName }}

	type {{.CName}} struct {
		{{.IName}}Impl
		*OTELTools
	}

	func New{{.CName}}(tools *OTELTools) *{{.CName}} {
		
		{{.IName}} := {{.IName}}Spanner{
			OTELTools: tools,
			next:  &{{.IName}}{},
		}
		return &{{.CName}}{
			{{.IName}}Impl: &{{.IName}},
		}
	}

	var _ {{.IName}}Impl = &{{ .IName }}{}

	type {{.IName}} struct {
	}
	{{range .IFuncs}} 
	    
		func (*{{ .IName }}) {{.FName}} ({{range .InParams}} {{.PName}} {{.PType}},{{end}}) ({{range .OutParams}}{{.PName}} {{.PType}},{{end}}) {
			//{{.FName}} business logic goes here
			return
		}
		{{end}}
	`

	spanTemplateStr :=
		`package {{ .PackageName }}

		var _ {{.IName}}Impl = &{{ .IName }}Spanner{}

		type {{ .IName }}Spanner struct {
			*OTELTools
			next {{ .IName }}Impl
		}
		{{range .IFuncs}} 
	    
		func (s *{{ .IName }}Spanner) {{.FName}} ({{range .InParams}} {{.PName}} {{.PType}},{{end}}) ({{range .OutParams}}{{.PName}} {{.PType}},{{end}}) {
			ctx, span := s.Tracer.Start(ctx, "{{.IName}}_{{.FName}}")
			ctx = ctx
			defer span.End()
			return s.next.{{.FName}}({{range .InParams}} {{.PName}},{{end}}) 
		}
		{{end}}

		`

	pkgName := "blogs"
	iName := "blogger"
	cName := strings.Title(iName)
	_ = IInfo{
		PackageName: pkgName,
		IName:       iName,
		CName:       cName,
		IFuncs: []IFunc{{IName: iName, FName: "addBlog", InParams: []FParam{{PName: "req", PType: "AddBlogRequest"}}, OutParams: []FParam{{PName: "res", PType: "AddBlogResponse"}}},
			{IName: iName, FName: "deleteBlog", InParams: []FParam{{PName: "req", PType: "DeleteBlogRequest"}}, OutParams: []FParam{{PName: "res", PType: "DeleteBlogResponse"}}},
			{IName: iName, FName: "listBlog", InParams: []FParam{{PName: "req", PType: "ListBlogRequest"}}, OutParams: []FParam{{PName: "res", PType: "ListBlogResponse"}}},
			{IName: iName, FName: "getBlog", InParams: []FParam{{PName: "req", PType: "GetBlogRequest"}}, OutParams: []FParam{{PName: "res", PType: "GetBlogResponse"}}}},
	}

	pkgName = "main"
	iName = "calculator"
	cName = strings.Title(iName)
	_ = IInfo{
		PackageName: pkgName,
		IName:       iName,
		CName:       cName,
		IFuncs: []IFunc{{IName: iName, FName: "add", InParams: []FParam{{PName: "x", PType: "int64"}, {PName: "y", PType: "int64"}}, OutParams: []FParam{{PName: "z", PType: "int64"}}},
			{IName: iName, FName: "sub", InParams: []FParam{{PName: "x", PType: "int64"}, {PName: "y", PType: "int64"}}, OutParams: []FParam{{PName: "z", PType: "int64"}}},
			{IName: iName, FName: "multiply", InParams: []FParam{{PName: "x", PType: "int64"}, {PName: "y", PType: "int64"}}, OutParams: []FParam{{PName: "z", PType: "int64"}}},
			{IName: iName, FName: "divide", InParams: []FParam{{PName: "x", PType: "int64"}, {PName: "y", PType: "int64"}}, OutParams: []FParam{{PName: "z", PType: "int64"}}}},
	}

	pkgName = "svc"
	iName = "userService"
	cName = strings.Title(iName)
	userDatabase := IInfo{
		PackageName: pkgName,
		IName:       iName,
		CName:       cName,
		IFuncs: []IFunc{
			{IName: iName, FName: "GetUsers", InParams: []FParam{{PName: "ctx", PType: "context.Context"}}, OutParams: []FParam{{PName: "users", PType: "[]*User"}, {PName: "err", PType: "error"}}},
			{IName: iName, FName: "GetUserByID", InParams: []FParam{{PName: "ctx", PType: "context.Context"}, {PName: "id", PType: "int64"}}, OutParams: []FParam{{PName: "user", PType: "*User"}, {PName: "err", PType: "error"}}},
			{IName: iName, FName: "CreateUser", InParams: []FParam{{PName: "ctx", PType: "context.Context"}, {PName: "req", PType: "CreateUserRequest"}}, OutParams: []FParam{{PName: "id", PType: "int64"}, {PName: "err", PType: "error"}}},
			{IName: iName, FName: "DeleteUser", InParams: []FParam{{PName: "ctx", PType: "context.Context"}, {PName: "req", PType: "DeleteUserRequest"}}, OutParams: []FParam{{PName: "err", PType: "error"}}},
			{IName: iName, FName: "UpdateUser", InParams: []FParam{{PName: "ctx", PType: "context.Context"}, {PName: "req", PType: "UpdateUserRequest"}}, OutParams: []FParam{{PName: "err", PType: "error"}}},
		}}
	doTemplate(userDatabase, interfaceStr, spanTemplateStr, starterStr)
	// doTemplate(calculator, interfaceStr, spanTemplateStr, starterStr)
	// doTemplate(crud, interfaceStr, spanTemplateStr, starterStr)
}

func doTemplate(calculator IInfo, interfaceStr string, spanTemplateStr string, starterStr string) {
	path := fmt.Sprintf("./%s/", calculator.PackageName)
	filename := fmt.Sprintf("%s%s.inf.go", path, calculator.PackageName)
	err := os.Mkdir(path, os.ModePerm)
	f, err := os.Create(filename)
	tmpl, err := template.New("test").Parse(interfaceStr)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = tmpl.Execute(f, calculator)
	if err != nil {
		panic(err)
	}
	filename = fmt.Sprintf("%s%s.span.go", path, calculator.PackageName)
	err = os.Mkdir(path, os.ModePerm)
	f, err = os.Create(filename)
	tmpl, err = template.New("test").Parse(spanTemplateStr)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = tmpl.Execute(f, calculator)
	if err != nil {
		panic(err)
	}
	filename = fmt.Sprintf("%s%s.go", path, calculator.PackageName)
	err = os.Mkdir(path, os.ModePerm)
	f, err = os.Create(filename)
	tmpl, err = template.New("test").Parse(starterStr)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = tmpl.Execute(f, calculator)
	if err != nil {
		panic(err)
	}
}
