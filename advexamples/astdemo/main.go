package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"github.com/kr/pretty"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"playgo/advexamples/astdemo/internal/visitors"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	os.Exit(realMain())
}
func realMain() (exitCode int) {
	app := cli.App("ast-demo", "demo for function type")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	app.Command("bs", "basic ", cli.ActionCommand(Basics))
	app.Command("sc", "scan ", cli.ActionCommand(Scan))
	app.Command("da", "dump ast ", cli.ActionCommand(DumpAST))
	app.Command("ps", "print source code ", cli.ActionCommand(PrintSourceCode))
	app.Command("na", "Navigating the AST ", cli.ActionCommand(NavigatingAst))
	app.Command("na2", "Navigating the AST: print like a tree ", cli.ActionCommand(NavigatingAst2))
	app.Command("vd", "tracking all short variable declarations ", cli.ActionCommand(TrackingVarDecl))
	app.Command("cs", "CollectStruts: 收集结构体定义 ", cli.ActionCommand(CollectStruts))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================

// ==============================================================================

func Basics() {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "test_data/simple.go", nil, 0)
	if err != nil {
		// Whoops!
		log.Fatal(err)
	}

	// 遍历是深度优先的
	ast.Walk(new(visitors.FuncVisitor), file)
	ast.Walk(new(visitors.ImportVisitor), file)

	printer.Fprint(os.Stdout, fset, file)

	// ------------------------
	// https://zupzup.org/go-ast-traversal/
	node := file
	fmt.Println("Imports:")
	for _, i := range node.Imports {
		fmt.Println(i.Path.Value)
	}
	fmt.Println("Comments:")
	for _, c := range node.Comments {
		fmt.Print(c.Text())
	}

	fmt.Println("Functions:")
	for _, f := range node.Decls {
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}
		fmt.Println(fn.Name.Name)
	}

	ast.Inspect(node, func(n ast.Node) bool {
		// Find Return Statements
		ret, ok := n.(*ast.ReturnStmt)
		if ok {
			fmt.Printf("return statement found on line %d:\n\t", fset.Position(ret.Pos()).Line)
			printer.Fprint(os.Stdout, fset, ret)
			return true
		}

		// Find Functions
		fn, ok := n.(*ast.FuncDecl)
		if ok {
			var exported string
			if fn.Name.IsExported() {
				exported = "exported "
			}
			fmt.Printf("%sfunction declaration found on line %d: \n\t%s\n", exported, fset.Position(fn.Pos()).Line, fn.Name.Name)
			return true
		}

		return true
	})

}

func Scan() {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", "package main; var a = 3", parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)
}

func DumpAST() {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", "package main; var a = 3", parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(f)
}

func PrintSourceCode() {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", "package main; var a = 3", parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	printer.Fprint(os.Stdout, fs, f)
}

// ----------
type visitor struct{}

func (v visitor) Visit(n ast.Node) ast.Visitor {
	fmt.Printf("%T\n", n)
	return v
}
func NavigatingAst() {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", "package main; var a = 3", parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

	var v visitor
	ast.Walk(v, f)
}
func NavigatingAst2() {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", "package main; var a = 3", parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

	var v visitor2
	ast.Walk(v, f)
}

type visitor2 int

func (v visitor2) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}
	fmt.Printf("%s%T\n", strings.Repeat("\t", int(v)), n)
	return v + 1
}

type visitor3 struct {
	locals map[string]int
}

func (v visitor3) Visit0(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}
	switch d := n.(type) {
	case *ast.AssignStmt:
		for _, name := range d.Lhs {
			if ident, ok := name.(*ast.Ident); ok {
				if ident.Name == "_" {
					continue
				}
				if ident.Obj != nil && ident.Obj.Pos() == ident.Pos() {
					v.locals[ident.Name]++
				}
			}
		}
	}
	return v
}

//
func (v visitor3) Visit(n ast.Node) ast.Visitor {
	// TODO 运行时有panic出现
	if n == nil {
		return nil
	}
	switch d := n.(type) {
	case *ast.AssignStmt:
		if d.Tok != token.DEFINE {
			return v
		}
		for _, name := range d.Lhs {
			v.local(name)
		}
	case *ast.RangeStmt:
		v.local(d.Key)
		v.local(d.Value)
	case *ast.FuncDecl:
		v.localList(d.Recv.List)
		v.localList(d.Type.Params.List)
		if d.Type.Results != nil {
			v.localList(d.Type.Results.List)
		}
	}
	return v
}
func (v visitor3) local(n ast.Node) {
	ident, ok := n.(*ast.Ident)
	if !ok {
		return
	}
	if ident.Name == "_" || ident.Name == "" {
		return
	}
	if ident.Obj != nil && ident.Obj.Pos() == ident.Pos() {
		v.locals[ident.Name]++
	}
}
func (v visitor3) localList(fs []*ast.Field) {
	for _, f := range fs {
		for _, name := range f.Names {
			v.local(name)
		}
	}
}

func TrackingVarDecl() {
	fs := token.NewFileSet()
	var v visitor3
	v = visitor3{
		locals: make(map[string]int),
	}

	file := "main.go"
	f, err := parser.ParseFile(fs, file, nil, parser.AllErrors)
	if err != nil {
		log.Printf("could not parse %s: %v", file, err)
		return
	}
	ast.Walk(v, f)
	fmt.Println(v.locals)

}

type visitor4 struct {
	locals  map[string]int
	globals map[string]int
}

func (v visitor4) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	switch d := n.(type) {
	case *ast.AssignStmt:
		if d.Tok != token.DEFINE {
			return v
		}
		fmt.Printf("%v\n", d.Lhs)
	case *ast.RangeStmt:
		fmt.Printf("%v, %v \n", d.Key, d.Value)

	}

	return v
}

func CollectStruts() {
	fs := token.NewFileSet()

	file := "main.go"
	f, err := parser.ParseFile(fs, file, nil, parser.AllErrors)
	if err != nil {
		log.Printf("could not parse %s: %v", file, err)
		return
	}

	s := collectStructs(f)
	pretty.Print(s)

}

/**
=========================================================================================
borrow from https://github.com/fatih/gomodifytags/blob/master/main.go
*/
// structType contains a structType node and it's name. It's a convenient
// helper type, because *ast.StructType doesn't contain the name of the struct
type structType struct {
	name string
	node *ast.StructType
}

// collectStructs collects and maps structType nodes to their positions
func collectStructs(node ast.Node) map[token.Pos]*structType {
	structs := make(map[token.Pos]*structType, 0)
	collectStructs := func(n ast.Node) bool {
		var t ast.Expr
		var structName string

		switch x := n.(type) {
		case *ast.TypeSpec:
			if x.Type == nil {
				return true

			}

			structName = x.Name.Name
			t = x.Type
		case *ast.CompositeLit:
			t = x.Type
		case *ast.ValueSpec:
			structName = x.Names[0].Name
			t = x.Type
		}

		x, ok := t.(*ast.StructType)
		if !ok {
			return true
		}

		structs[x.Pos()] = &structType{
			name: structName,
			node: x,
		}
		return true
	}
	ast.Inspect(node, collectStructs)
	return structs
}

// ===================================
/**
https://github.com/campoy/justforfunc/blob/master/25-go-parser/main.go
https://github.com/fatih/gomodifytags/blob/master/main.go
*/
