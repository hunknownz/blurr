//go:build ignorepackage

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	typeNames = flag.String("type", "", "comma-separated list of int type names; must be set")
)

const (
	dtypeFuncTmpl = `
func (dt %s) DType() DType {
    return D%s
}
`
	itemSizeFuncTmpl = `
func (dt %s) ItemSize() int {
	return int(unsafe.Sizeof(dt))
}
`
)

type Generator struct {
	buf bytes.Buffer
}

func (g *Generator) printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

func (g *Generator) format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	return src
}

func main() {
	flag.Parse()
	if len(*typeNames) == 0 {
		log.Fatal("-type must be set")
	}
	pkgName, fileName := os.Getenv("GOPACKAGE"), os.Getenv("GOFILE")

	gen := new(Generator)
	gen.printf("// Code generated by \"go run gen_dtype.go\"; DO NOT EDIT.\n")
	gen.printf("package %s", pkgName)
	gen.printf("\n")
	gen.printf("import \"unsafe\"\n")

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fileName, nil, 0)
	if err != nil {
		log.Fatal(err)
	}
	ast.Inspect(f, func(n ast.Node) bool {
		decl, ok := n.(*ast.GenDecl)
		if !ok || decl.Tok != token.TYPE {
			return true
		}
		for _, spec := range decl.Specs {
			vspec := spec.(*ast.TypeSpec)

			dttype := vspec.Name.Name
			identType, ok := vspec.Type.(*ast.Ident)
			if !ok {
				continue
			}
			dtype := identType.Name
			dtype = strings.ToUpper(dtype[:1]) + dtype[1:]
			gen.printf(dtypeFuncTmpl, dttype, dtype)
			gen.printf("\n")
			gen.printf(itemSizeFuncTmpl, dttype)
		}
		return true
	})

	src := gen.format()
	fNames := strings.Split(fileName, ".")
	baseName := fmt.Sprintf("%s_gen.go", fNames[0])
	outputName := filepath.Join(".", strings.ToLower(baseName))
	err = ioutil.WriteFile(outputName, src, 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}
}
