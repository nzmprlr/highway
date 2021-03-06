package env

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	now := time.Now().UTC().String()
	tests := []struct {
		key string
		def string
		get string
		set string
	}{
		{
			key: now,
			def: now,
			get: now,
			set: "test",
		}, {
			key: now,
			def: now,
			get: "test",
		}, {
			key: now,
			def: now,
			get: now,
		},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.key, func(t *testing.T) {
			get := Get(tc.key, tc.def)

			assert.Equal(t, tc.get, get)
			if tc.set == "" {
				assert.NoError(t, os.Unsetenv(tc.key))
			} else {
				assert.NoError(t, os.Setenv(tc.key, tc.set))
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	now := time.Now().UTC().String()
	unix := int(time.Now().Unix())
	tests := []struct {
		key string
		def int
		get int
		set string
	}{
		{
			key: now,
			def: unix,
			get: unix,
			set: "test",
		}, {
			key: now,
			def: unix,
			get: unix,
			set: "-1",
		}, {
			key: now,
			def: unix,
			get: -1,
		}, {
			key: now,
			def: unix,
			get: unix,
			set: "1",
		}, {
			key: now,
			def: unix,
			get: 1,
			set: "2147483648", // overflow int
		}, {
			key: now,
			def: unix,
			get: unix,
		},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.key, func(t *testing.T) {
			get := GetInt(tc.key, tc.def)

			assert.Equal(t, tc.get, get)
			if tc.set == "" {
				assert.NoError(t, os.Unsetenv(tc.key))
			} else {
				assert.NoError(t, os.Setenv(tc.key, tc.set))
			}
		})
	}
}
