package engine

import (
	"github.com/yaz3/achievement-system/pgk/evaluationManager"

	"github.com/PaesslerAG/gval"
	//"github.com/achievement-system/evaluationManager"
)

type Expression struct {
	ExpressionToEvaluate string
	Typ                  int
}

const (
	Integer int = iota
	Float
	Boolean
	String
	StringSlices
)

func AddNewExpression(expression Expression) error {
	switch expression.Typ {
	case Integer:
		eval, err := gval.Full().NewEvaluable("1 == x")
		evaluationManager.GetManager().AddExp(expression.ExpressionToEvaluate, eval)

		//func (e Evaluable) EvalInt(c context.Context, parameter interface{}) (int, error)
	}
	return nil
}
