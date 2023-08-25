package valconv_test

import (
	. "cbsutil/valconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type typeBool bool

func TestValueAny_Bool(t *testing.T) {
	v1 := Any(true).Bool().Value
	assert.True(t, v1)
	v2 := Any("ab").Bool().Err
	assert.Error(t, v2)

	v3 := Any(true).AsBool().Value
	assert.True(t, v3)
	v4 := Any(typeBool(true)).AsBool().Value
	assert.True(t, v4)

	v5 := Any("true").ToBool().Value
	assert.True(t, v5)

	v6 := Any("@#!@#@#").ToBool().Err
	t.Log(v6)
	assert.Error(t, v6)
}
