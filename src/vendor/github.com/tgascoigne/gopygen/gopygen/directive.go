package gopygen

import (
	"go/ast"
	"go/token"
	"regexp"
	"strings"
)

var directiveRegexp = regexp.MustCompile(`gopygen:\s?"(.*?)"`)

type DirectiveData struct {
	Tags []string
}

type Directive struct {
	*DirectiveData
	fileset *token.FileSet
}

func NewDirective(fileset *token.FileSet) Directive {
	return Directive{
		fileset: fileset,
		DirectiveData: &DirectiveData{
			Tags: []string{},
		},
	}
}

func (t Directive) Visit(n ast.Node) ast.Visitor {
	switch node := n.(type) {
	case *ast.Comment:
		submatches := directiveRegexp.FindStringSubmatch(node.Text)

		if len(submatches) == 0 {
			return nil
		}

		t.Tags = append(t.Tags, submatches[1])
		return nil
	}
	return t
}

func (t *DirectiveData) String() string {
	return strings.Join(t.Tags, " ")
}
