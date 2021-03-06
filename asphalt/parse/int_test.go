package parse

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt8(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in  []string
		out []int8
		err bool
	}{
		{
			in:  []string{""},
			err: true,
		}, {
			in:  []string{"0", ""},
			err: true,
		}, {
			in:  []string{"-129"},
			err: true,
		}, {
			in:  []string{"128"},
			err: true,
		}, {
			in:  []string{"-128", "-5", "0", "5", "127"},
			out: []int8{-128, -5, 0, 5, 127},
		},
	}

	for _, test := range tests {
		tc := test
		t.Run(fmt.Sprintf("%s", tc.in), func(t *testing.T) {
			p, err := Int8(tc.in[0], tc.in[1:]...)

			if tc.err {
				assert.Error(t, err)
			} else {
				for i, pi := range p {
					assert.Equal(t, tc.out[i], pi)
				}
			}
		})
	}
}

func TestInt16(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in  []string
		out []int16
		err bool
	}{
		{
			in:  []string{""},
			err: true,
		}, {
			in:  []string{"0", ""},
			err: true,
		}, {
			in:  []string{"-32769"},
			err: true,
		}, {
			in:  []string{"32768"},
			err: true,
		}, {
			in:  []string{"-32768", "-555", "0", "555", "32767"},
			out: []int16{-32768, -555, 0, 555, 32767},
		},
	}

	for _, test := range tests {
		tc := test
		t.Run(fmt.Sprintf("%s", tc.in), func(t *testing.T) {
			p, err := Int16(tc.in[0], tc.in[1:]...)

			if tc.err {
				assert.Error(t, err)
			} else {
				for i, pi := range p {
					assert.Equal(t, tc.out[i], pi)
				}
			}
		})
	}
}

func TestInt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in  []string
		out []int
		err bool
	}{
		{
			in:  []string{""},
			err: true,
		}, {
			in:  []string{"0", ""},
			err: true,
		}, {
			in:  []string{"-2147483649"},
			err: true,
		}, {
			in:  []string{"2147483648"},
			err: true,
		}, {
			in:  []string{"-2147483648", "-555", "0", "555", "2147483647"},
			out: []int{-2147483648, -555, 0, 555, 2147483647},
		},
	}

	for _, test := range tests {
		tc := test
		t.Run(fmt.Sprintf("%s", tc.in), func(t *testing.T) {
			p, err := Int(tc.in[0], tc.in[1:]...)
			p32, err32 := Int32(tc.in[0], tc.in[1:]...)

			if tc.err {
				assert.Error(t, err)
				assert.Error(t, err32)
			} else {
				for i, pi := range p {
					assert.Equal(t, tc.out[i], pi)
				}
				for i, pi := range p32 {
					assert.Equal(t, int32(tc.out[i]), pi)
				}

			}
		})
	}
}

func TestInt64(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in  []string
		out []int64
		err bool
	}{
		{
			in:  []string{""},
			err: true,
		}, {
			in:  []string{"0", ""},
			err: true,
		}, {
			in:  []string{"–9223372036854775809"},
			err: true,
		}, {
			in:  []string{"9223372036854775808"},
			err: true,
		}, {
			in:  []string{"-9223372036854775808", "-5555555555555", "0", "5555555555555", "9223372036854775807"},
			out: []int64{-9223372036854775808, -5555555555555, 0, 5555555555555, 9223372036854775807},
		},
	}

	for _, test := range tests {
		tc := test
		t.Run(fmt.Sprintf("%s", tc.in), func(t *testing.T) {
			p, err := Int64(tc.in[0], tc.in[1:]...)

			if tc.err {
				assert.Error(t, err)
			} else {
				for i, pi := range p {
					assert.Equal(t, tc.out[i], pi)
				}
			}
		})
	}
}
