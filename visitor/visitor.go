package visitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/merideum/merideum-go/parser"
)

type MerideumMeritVisitor struct {
	antlr.ParseTreeVisitor
	Variables map[string]MeritVariable
}

type MeritVariable struct {
	Name        string
	Value       any
	Initialized bool
	// ValueType VariableType
}

type VoidMeritValue struct{}

type MeritValue struct {
	Value any
}

func NewMeritValue(value any) MeritValue {
	return MeritValue{
		Value: value,
	}
}

func NewMerideumMeritVisitor() parser.MeritVisitor {
	return &MerideumMeritVisitor{
		ParseTreeVisitor: &parser.BaseMeritVisitor{},
		Variables:        make(map[string]MeritVariable),
	}
}

func (v *MerideumMeritVisitor) VisitIntegerExpression(ctx *parser.IntegerExpressionContext) interface{} {
	return NewMeritValue(ctx.GetText())
}

func (v *MerideumMeritVisitor) VisitAssignment(ctx *parser.AssignmentContext) interface{} {
	return NewMeritValue(v.Visit(ctx.Expression()).(MeritValue).Value)
}

func (v *MerideumMeritVisitor) VisitVariableModifier(ctx *parser.VariableModifierContext) interface{} {
	return NewMeritValue(ctx.GetText())
}

func (v *MerideumMeritVisitor) VisitVariableAssignment(ctx *parser.VariableAssignmentContext) interface{} {
	initialized := false

	name := ctx.IDENTIFIER().GetText()

	var value MeritValue = MeritValue{
		Value: nil,
	}

	if ctx.Assignment() != nil {
		value = v.Visit(ctx.Assignment()).(MeritValue)
		initialized = true
	}

	v.Variables[name] = MeritVariable{
		Name:        name,
		Value:       value,
		Initialized: initialized,
	}

	return nil
}

func (v *MerideumMeritVisitor) VisitOutputAssignment(ctx *parser.OutputAssignmentContext) interface{} {
	return nil
}

func (v *MerideumMeritVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MerideumMeritVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MerideumMeritVisitor) VisitParse(ctx *parser.ParseContext) interface{} {
	return v.VisitChildren(ctx)
}
