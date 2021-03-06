package highway

import "github.com/nzmprlr/highway/toll"

func API(t *toll.Toll, check interface{}, newFn func(t *toll.Toll) interface{}) interface{} {
	if check == nil {
		return newFn(t)
	}

	return check
}
