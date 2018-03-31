package createarray

import (
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("create-array")

type CreateArray struct {
}

func init() {
	function.Registry(&CreateArray{})
}

func (s *CreateArray) GetName() string {
	return "createArray"
}

func (s *CreateArray) GetCategory() string {
	return "object"
}

func (s *CreateArray) Eval(objs ...interface{}) ([]interface{}, error) {
	log.Debugf("Start create array function with parameters %s", objs)
	var objArray []interface{}

	for _, obj := range objs {
		switch t := obj.(type) {
		case []interface{}:
			objArray = append(objArray, t)
		case interface{}:
			v, err := data.CoerceToObject(t)
			if err != nil {
				objArray = append(objArray, t)
			} else {
				objArray = append(objArray, v)
			}
		}

	}
	return objArray, nil
}
