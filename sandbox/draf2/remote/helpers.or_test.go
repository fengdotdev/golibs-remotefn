package remote_test

import (
	"testing"

	"github.com/fengdotdev/golibs-remotefn/sandbox/draf2/remote"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestOr(t *testing.T) {

	m := map[string]interface{}{
		"key1": "value1",
		"key2": 42,
	}

	valueString := remote.Or(m, "key1", "default")
	assert.EqualWithMessage(t, valueString, "value1", "Expected value for key1 to be 'value1'")

	valueInt := remote.Or(m, "key2", 0)
	assert.EqualWithMessage(t, valueInt, 42, "Expected value for key2 to be 42")
	valueDefault := remote.Or(m, "key3", "default")
	assert.EqualWithMessage(t, valueDefault, "default", "Expected default value for key3 to be 'default'")
	valueIntDefault := remote.Or(m, "key3", 0)
	assert.EqualWithMessage(t, valueIntDefault, 0, "Expected default value for key3 to be 0")

}

func TestOrErr(t *testing.T) {
	m := map[string]interface{}{
		"key1": "value1",
		"key2": 42,
	}

	valueString, err := remote.OrErr[string](m, "key1")
	assert.NoErrorWithMessage(t, err, "Expected no error for key1")
	assert.EqualWithMessage(t, valueString, "value1", "Expected value for key1 to be 'value1'")

	valueInt, err := remote.OrErr[int](m, "key2")
	assert.NoErrorWithMessage(t, err, "Expected no error for key2")
	assert.EqualWithMessage(t, valueInt, 42, "Expected value for key2 to be 42")
	_, err = remote.OrErr[string](m, "key3")
	assert.ErrorWithMessage(t, err, "Expected error for key3 not found")


	m2 := remote.ResultSingle(5)
	valueInt2, err := remote.OrErr[int](m2, "result")
	assert.NoErrorWithMessage(t, err, "Expected no error for key 'result'")
	assert.EqualWithMessage(t, valueInt2, 5, "Expected value for key 'result' to be 5")
}
