package search

import (
	"fmt"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/ref"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("search-object")

type Search struct {
}

func init() {
	function.Registry(&Search{})
}

func (s *Search) GetName() string {
	return "search"
}

func (s *Search) GetCategory() string {
	return "object"
}

func (s *Search) Eval(obj interface{}, mapRef string) (interface{}, error) {

	if obj == nil {
		return nil, nil
	} else {
		v, err := data.CoerceToObject(obj)
		if err != nil {
			return nil, fmt.Errorf("Object search must be a json object")
		}
		obj = v
	}

	if strings.HasPrefix(mapRef, "$.") {
		arrayRef := ref.NewArrayRef(mapRef)
		return arrayRef.EvalFromData(obj)
	}

	return nil, fmt.Errorf("Search object map ref must start with %.")

}
