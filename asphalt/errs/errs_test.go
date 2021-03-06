package errs

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewErrs(t *testing.T) {
	t.Parallel()

	const addedError = "added error"
	err := New()

	tests := []struct {
		err     *Errs
		nilFn   assert.ValueAssertionFunc
		e1Error string
		errorFn assert.ErrorAssertionFunc
		e2Error string
	}{
		{
			err:     nil,
			nilFn:   assert.Nil,
			e1Error: emptyErrorMsg,
			errorFn: assert.NoError,
			e2Error: addedError,
		}, {
			err:     &Errs{},
			nilFn:   assert.NotNil,
			e1Error: emptyErrorMsg,
			errorFn: assert.Error,
			e2Error: emptyErrorMsg + "\n" + addedError,
		}, {
			err:     err, // New()
			nilFn:   assert.NotNil,
			e1Error: emptyErrorMsg,
			errorFn: assert.Error,
			e2Error: emptyErrorMsg + "\n" + addedError,
		},
	}

	for i, test := range tests {
		tc := test
		t.Run(strconv.FormatInt(int64(i), 10), func(t *testing.T) {
			tc.nilFn(t, tc.err)
			assert.Empty(t, tc.err.Error())
			assert.EqualError(t, tc.err, "")
			tc.err.Add(nil)
			assert.Nil(t, tc.err.Reduce())
			assert.EqualError(t, tc.err.Add(errors.New("")), tc.e1Error)
			tc.errorFn(t, tc.err.Reduce())
			assert.EqualError(t, tc.err.Add(errors.New(addedError)), tc.e2Error)
			tc.errorFn(t, tc.err.Reduce())
		})
	}

	err1 := New()
	err1.Add(errors.New(addedError))

	err.Add(err1)
	assert.Error(t, err1)
	assert.EqualError(t, err.Reduce(), emptyErrorMsg+"\n"+addedError+"\n"+addedError)
	assert.EqualError(t, err, emptyErrorMsg+"\n"+addedError+"\n"+addedError)
}
