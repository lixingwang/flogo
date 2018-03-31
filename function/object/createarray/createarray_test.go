package createarray

import (
	"testing"

	"fmt"

	"encoding/json"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/stretchr/testify/assert"
)

var s = &CreateArray{}

func TestCreateArray_Eval(t *testing.T) {
	v, err := s.Eval(`{"name":"tracy"}`, `{"name":"li"}`)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v[1].(map[string]interface{})["name"])
}

func TestExpression(t *testing.T) {
	fun, err := expression.NewFunctionExpression(`object.createArray("{\"name\":\"tracy\"}", "{\"name\":\"li\"}")`).GetFunction()
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	ret := v[0].([]interface{})
	assert.Equal(t, "tracy", ret[0].(map[string]interface{})["name"])
	v2, _ := json.Marshal(v[0])
	fmt.Println(string(v2))
}

func TestExpression2(t *testing.T) {
	fun, err := expression.NewFunctionExpression(`object.createArray("name", "value")`).GetFunction()
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	//ret := v[0].([]interface{})
	v2, _ := json.Marshal(v[0])
	fmt.Println(string(v2))
}
