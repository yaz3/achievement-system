package evaluationmanager

import (
	"fmt"
	"sync"

	"github.com/PaesslerAG/gval"
)

//expressionManager
type expressionsManager struct {
	evaluations map[string]gval.Evaluable
}

var singleton *expressionsManager
var once sync.Once

//GetManager Get ExpressionManager singleton
func GetManager() *expressionsManager {
	once.Do(func() {
		singleton = &expressionsManager{evaluations: make(map[string]gval.Evaluable)}
	})
	return singleton
}

func (evalM *expressionsManager) AddExp(name string, exp gval.Evaluable) error {
	if _, ok := evalM.evaluations[name]; ok {
		return fmt.Errorf("An Expression under the name: %s already exists", name)
	}
	evalM.evaluations[name] = exp
	return nil
}
