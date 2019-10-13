package engine

import (
	//"github.com/yaz3/achievement-system/pgk/evaluationManager"

	//"achievements-system/pkg/evaluationManager"
	"fmt"
	"log"
	"sync"

	"context"

	"github.com/PaesslerAG/gval"
	//"github.com/achievement-system/evaluationManager"
)

const (
	Integer int = iota
	Float
	Boolean
	String
	StringSlices
)

//Expression ..
type Expression struct {
	AchievementName      string
	ExpressionToEvaluate string
	Typ                  int
	Evaluable            gval.Evaluable
}

//ExpressionManager ...
type ExpressionsManager struct {
	evaluations map[string]Expression
}

var singleton *ExpressionsManager
var once sync.Once

func initSingleton() *ExpressionsManager {
	singleton = &ExpressionsManager{evaluations: make(map[string]Expression)}
	//TODO: InitFromFile (path parameter create expressions and populate map)
	return singleton

}

//GetManager Get ExpressionManager singleton
func GetManager() *ExpressionsManager {
	once.Do(func() {
		singleton = initSingleton()
	})
	return singleton
}

//AddExp ...
func (evalM *ExpressionsManager) AddExp(name string, exp Expression) error {
	if _, ok := evalM.evaluations[name]; ok {
		return fmt.Errorf("An Expression under the name: %s already exists", name)
	}
	evalM.evaluations[name] = exp
	return nil
}

//CreateNewExpressionEvaluable ...
func (expression Expression) CreateNewExpressionEvaluable() error {
	eval, err := gval.Full().NewEvaluable(expression.ExpressionToEvaluate)
	if err != nil {
		return fmt.Errorf("expression %s coudln't be created", expression.AchievementName)
	}
	expression.Evaluable = eval
	GetManager().AddExp(expression.AchievementName, expression)

	return nil
}

//BatchAchievements ...
func (evalM *ExpressionsManager) BatchAchievements(data map[string]interface{}) map[string]interface{} {
	if len(evalM.evaluations) < 1 {
		return nil
	}
	var results = make(map[string]interface{})

	for achievement, expression := range evalM.evaluations {
		//expression.
		log.Println(expression, achievement)
		result, err := expression.ExecuteExpression(data)
		if err != nil {
			//TODO: DEfine the best behavior for this case
			continue
		}
		results[expression.AchievementName] = result
	}
	return nil
}

func ConvertTyp(typ int, data interface{}) interface{} {
	switch typ {
	case Integer:
		return data.(int)
	case String:
		return data.(string)
	case Float:
		return data.(float64)
	case Boolean:
		return data.(bool)
	default:
		return nil
	}

}

//ExecuteExpression ...
func (expression Expression) ExecuteExpression(data map[string]interface{}) (interface{}, error) {
	background := context.Background()
	switch expression.Typ {
	case Integer:
		res, err := expression.Evaluable.EvalInt(background, data)
		return res, err
	case String:
		res, err := expression.Evaluable.EvalString(background, data)
		return res, err
	case Float:
		res, err := expression.Evaluable.EvalFloat64(background, data)
		return res, err
	case Boolean:
		res, err := expression.Evaluable.EvalBool(background, data)
		return res, err
	default:
		return nil, fmt.Errorf("Unknown type for expression %s", expression.AchievementName)
	}
}
