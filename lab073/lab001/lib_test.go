package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Student struct {
	Sid   int
	Sname string
}

func TestSomething(t *testing.T) {
	//assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	//assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	//assert for nil
	var s *Student
	assert.Nil(t, s)

	if assert.NotNil(t, s) {

	}
}
