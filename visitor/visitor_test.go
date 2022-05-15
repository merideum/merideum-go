package visitor

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/merideum/merideum-go/parser"
)

func TestVisitorHelloWorld(t *testing.T) {
	expression := "const test = 123"
	input := antlr.NewInputStream(expression)
	lexer := parser.NewMeritLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewMeritParser(stream)
	parser.BuildParseTrees = true
	tree := parser.Parse()
	visitor := NewMerideumMeritVisitor().(*MerideumMeritVisitor)

	visitor.Visit(tree)

	var element any
	var present bool

	element, present = visitor.Variables["test"]

	if present {
		castElement, present := element.(int32)
		if !present {
			t.Errorf("Expected 123 but got %q", castElement)
		}
	} else {
		t.Error("Missing element 'test'")
	}
}
