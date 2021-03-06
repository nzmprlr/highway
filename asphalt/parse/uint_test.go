package parse

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUint8(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in  []string
		out []uint8
		err bool
	}{
		{
			in:  []string{""},
			err: true,
		}, {
			in:  []string{"0", ""},
			err: true,
		}, {
			in:  []string{"-1"},
			err: true,
		}, {
			in:  []string{"256"},
			err: true,
		}, {
			in:  []string{"0", "5", "255"},
			out: []uint8{0, 5, 255},
		},
	}

	for _, test := range tests {
		tc := test
		t.Run(fmt.Sprintf("%s", tc.in), func(t *testing.T) {
			p, err := Uint8(tc.in[0], tc.in[1:]...)

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

func TestUint16(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in  []string
		out []uint16
		err bool
	}{
		{
			in:  []string{""},
			err: true,
		}, {
			in:  []string{"0", ""},
			err: true,
		}, {
			in:  []string{"-1"},
			err: true,
		}, {
			in:  []string{"65536"},
			err: true,
		}, {
			in:  []string{"0", "5555", "65535"},
			out: []uint16{0, 5555, 65535},
		},
	}

	for _, test := range tests {
		tc := test
		t.Run(fmt.Sprintf("%s", tc.in), func(t *testing.T) {
			p, err := Uint16(tc.in[0], tc.in[1:]...)

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

func TestUint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in  []string
		out []uint
		err bool
	}{
		{
			in:  []string{""},
			err: true,
		}, {
			in:  []string{"0", ""},
			err: true,
		}, {
			in:  []string{"-1"},
			err: true,
		}, {
			in:  []string{"4294967296"},
			err: true,
		}, {
			in:  []string{"0", "555", "4294967295"},
			out: []uint{0, 555, 4294967295},
		},
	}

	for _, test := range tests {
		tc := test
		t.Run(fmt.Sprintf("%s", tc.in), func(t *testing.T) {
			p, err := Uint(tc.in[0], tc.in[1:]...)
			p32, err32 := Uint32(tc.in[0], tc.in[1:]...)

			if tc.err {
				assert.Error(t, err)
				assert.Error(t, err32)
			} else {
				for i, pi := range p {
					assert.Equal(t, tc.out[i], pi)
				}
				for i, pi := range p32 {
					assert.Equal(t, uint32(tc.out[i]), pi)
				}
			}
		})
	}
}

func TestUint64(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in  []string
		out []uint64
		err bool
	}{
		{
			in:  []string{""},
			err: true,
		}, {
			in:  []string{"0", ""},
			err: true,
		}, {
			in:  []string{"-1"},
			err: true,
		}, {
			in:  []string{"18446744073709551616"},
			err: true,
		}, {
			in:  []string{"0", "5555555555555", "18446744073709551615"},
			out: []uint64{0, 5555555555555, 18446744073709551615},
		},
	}

	for _, test := range tests {
		tc := test
		t.Run(fmt.Sprintf("%s", tc.in), func(t *testing.T) {
			p, err := Uint64(tc.in[0], tc.in[1:]...)

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
