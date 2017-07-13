package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloatEqual(t *testing.T) {
	//123.45 123.456 不相等
	assert.NotEqual(t, 123.45, 123.456)

	//123.45 123.450 相等
	assert.Equal(t, 123.45, 123.450)

	//InEpsilon 相对误差
	assert.InEpsilon(t, 123.4567, 123.4566, 0.0001)
	assert.InEpsilon(t, 123.456, 123.458, 0.0001)
	assert.InEpsilon(t, 123.456, 123.458, 0.1)

	//delta 值在+-delta之间就可以
	assert.InDelta(t, 123.4567, 123.4566, 0.001)
	assert.InDelta(t, 123.4577, 123.4566, 0.001)
}
