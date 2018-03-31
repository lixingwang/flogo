package getelement

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("get-element")

type GetElement struct {
}

func init() {
	function.Registry(&GetElement{})
}

func (s *GetElement) GetName() string {
	return "getElement"
}

func (s *GetElement) GetCategory() string {
	return "array"
}

func (s *GetElement) Eval(arrys []interface{}, index int) interface{} {
	log.Debugf("Get array element  %+v by index %d ", arrys, index)
	if arrys == nil || len(arrys) <= 0 {
		return nil
	}
	return arrys[index]
}
