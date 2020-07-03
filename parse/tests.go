package parse

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"github.com/pkg/errors"
)

// functions returns an array of function names found in file or src
func parse(file string, src interface{}) ([]string, error) {
	names := make([]string, 0)
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file, src, parser.AllErrors)
	if err != nil {
		return names, errors.Wrap(err, "parser.parsefile")
	}
	for _, d := range f.Decls {
		if fn, isFn := d.(*ast.FuncDecl); isFn {
			names = append(names, fn.Name.String())
		}
	}
	return names, nil
}

// tests parses the file and returns the Test functions
func Tests(file string, src interface{}) ([]string, error) {
	names, err := parse(file, src)
	if err != nil {
		return names, errors.Wrap(err, "parse")
	}
	new := make([]string, 0, len(names))
	for _, s := range names {
		if strings.HasPrefix(s, "Test") {
			new = append(new, s)
		}
	}
	return new, nil
}
