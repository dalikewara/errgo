package errgo_test

import (
	"github.com/dalikewara/errgo"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNew test New function.
func TestNew(t *testing.T) {
	err := errgo.New("01", "data not found")
	assert.NotNil(t, err)
	assert.Implements(t, (*errgo.ErrGo)(nil), err)
}

// TestNewWithHttpStatus test NewWithHttpStatus function.
func TestNewWithHttpStatus(t *testing.T) {
	err := errgo.NewWithHttpStatus("01", "data not found", 200)
	assert.NotNil(t, err)
	assert.Implements(t, (*errgo.ErrGo)(nil), err)
}

// TestErrgo_GetCode test errgo.GetCode method.
func TestErrgo_GetCode(t *testing.T) {
	err := errgo.New("01", "data not found")
	assert.NotNil(t, err)
	assert.Equal(t, "01", err.GetCode())
}

// TestErrgo_GetMessage test errgo.GetMessage method.
func TestErrgo_GetMessage(t *testing.T) {
	err := errgo.New("01", "data not found")
	assert.NotNil(t, err)
	assert.Equal(t, "data not found", err.GetMessage())
}

// TestErrgo_GetHttpStatus test errgo.GetHttpStatus method.
func TestErrgo_GetHttpStatus(t *testing.T) {
	t.Run("http status 0", func(t *testing.T) {
		err := errgo.New("01", "data not found")
		assert.NotNil(t, err)
		assert.Equal(t, 0, err.GetHttpStatus())
	})
	t.Run("OK", func(t *testing.T) {
		err := errgo.NewWithHttpStatus("01", "data not found", 200)
		assert.NotNil(t, err)
		assert.Equal(t, 200, err.GetHttpStatus())
		assert.Equal(t, "data not found", err.GetMessage())
		assert.Equal(t, "01", err.GetCode())
		assert.EqualError(t, err.GetError(), "01||data not found")
	})
}

// TestErrgo_GetError test errgo.GetError method.
func TestErrgo_GetError(t *testing.T) {
	err := errgo.New("01", "data not found")
	assert.NotNil(t, err)
	assert.EqualError(t, err.GetError(), "01||data not found")
}
