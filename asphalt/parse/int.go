package parse

import (
	"strconv"

	"github.com/nzmprlr/highway/asphalt/errs"
)

func parseInt(s string, size int) (int64, error) {
	return strconv.ParseInt(s, 10, size)
}

func Int8(i string, s ...string) ([]int8, error) {
	errs := errs.New()

	s = append([]string{i}, s...)
	p := make([]int8, len(s))
	for i, v := range s {
		c, err := parseInt(v, 8)
		if err != nil {
			errs.Add(err)
		}

		p[i] = int8(c)
	}

	return p, errs.Reduce()
}

func Int16(i string, s ...string) ([]int16, error) {
	errs := errs.New()

	s = append([]string{i}, s...)
	p := make([]int16, len(s))
	for i, v := range s {
		c, err := parseInt(v, 16)
		if err != nil {
			errs.Add(err)
		}

		p[i] = int16(c)
	}

	return p, errs.Reduce()
}

func Int(i string, s ...string) ([]int, error) {
	errs := errs.New()

	s = append([]string{i}, s...)
	p := make([]int, len(s))
	for i, v := range s {
		c, err := parseInt(v, 32)
		if err != nil {
			errs.Add(err)
		}

		p[i] = int(c)
	}

	return p, errs.Reduce()
}

func Int32(i string, s ...string) ([]int32, error) {
	p, err := Int(i, s...)
	if err != nil {
		return nil, err
	}

	r := make([]int32, len(p))
	for i, v := range p {
		r[i] = int32(v)
	}

	return r, nil
}

func Int64(i string, s ...string) ([]int64, error) {
	errs := errs.New()

	s = append([]string{i}, s...)
	p := make([]int64, len(s))
	for i, v := range s {
		c, err := parseInt(v, 64)
		if err != nil {
			errs.Add(err)
		}

		p[i] = int64(c)
	}

	return p, errs.Reduce()
}
