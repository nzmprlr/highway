package parse

import (
	"strconv"

	"github.com/nzmprlr/highway/asphalt/errs"
)

func parseUint(s string, size int) (uint64, error) {
	return strconv.ParseUint(s, 10, size)
}

// Uint8 parses string arguments to uint8 array by args order.
func Uint8(i string, s ...string) ([]uint8, error) {
	errs := errs.New()

	s = append([]string{i}, s...)
	p := make([]uint8, len(s))
	for i, v := range s {
		c, err := parseUint(v, 8)
		if err != nil {
			errs.Add(err)
		}

		p[i] = uint8(c)
	}

	return p, errs.Reduce()
}

func Uint16(i string, s ...string) ([]uint16, error) {
	errs := errs.New()

	s = append([]string{i}, s...)
	p := make([]uint16, len(s))
	for i, v := range s {
		c, err := parseUint(v, 16)
		if err != nil {
			errs.Add(err)
		}

		p[i] = uint16(c)
	}

	return p, errs.Reduce()
}

func Uint(i string, s ...string) ([]uint, error) {
	errs := errs.New()

	s = append([]string{i}, s...)
	p := make([]uint, len(s))
	for i, v := range s {
		c, err := parseUint(v, 32)
		if err != nil {
			errs.Add(err)
		}

		p[i] = uint(c)
	}

	return p, errs.Reduce()
}

func Uint32(i string, s ...string) ([]uint32, error) {
	p, err := Uint(i, s...)
	if err != nil {
		return nil, err
	}

	r := make([]uint32, len(p))
	for i, v := range p {
		r[i] = uint32(v)
	}

	return r, nil
}

func Uint64(i string, s ...string) ([]uint64, error) {
	errs := errs.New()

	s = append([]string{i}, s...)
	p := make([]uint64, len(s))
	for i, v := range s {
		c, err := parseUint(v, 64)
		if err != nil {
			errs.Add(err)
		}

		p[i] = uint64(c)
	}

	return p, errs.Reduce()
}
