package search

import (
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/stretchr/testify/assert"
)

var s = &Search{}

func TestCreateArray_Eval(t *testing.T) {
	v, err := s.Eval(`{"name":"tracy"}`, `$.name`)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	assert.Equal(t, "tracy", v)
}

func TestExpression(t *testing.T) {
	fun, err := expression.NewFunctionExpression(`object.search("{\"name\":\"tracy\"}", $.name)`).GetFunction()
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.Equal(t, "tracy", v[0])
}
