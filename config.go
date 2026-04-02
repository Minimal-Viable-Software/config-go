/*
Package config implements a minimal viable config library.

Only configuration through environment variables are supported.
*/
package config

import (
	"encoding"
	"flag"
	"fmt"
	"os"
	"strings"
)

var prefix = ""

// SetPrefix adds a prefix to all keys used to set values.
func SetPrefix(s string) {
	prefix = strings.ToLower(s)
}

func setValue(value flag.Value, key string) error {
	key = fmt.Sprintf("%s%s", prefix, strings.ToLower(key))

	for _, pair := range os.Environ() {
		envkey, envval, ok := strings.Cut(pair, "=")
		if !ok || key != strings.ToLower(envkey) {
			continue
		}
		if err := value.Set(envval); err != nil {
			return fmt.Errorf("%s: bad value: %w", key, err)
		}
	}

	return nil
}

// String sets a string value from an environment variable.
func String(p *string, key string) error {
	return setValue((*stringValue)(p), key)
}

// Int sets an int value from an environment variable.
func Int(p *int, key string) error {
	return setValue((*intValue)(p), key)
}

// Int64 sets an int64 value from an environment variable.
func Int64(p *int64, key string) error {
	return setValue((*int64Value)(p), key)
}

// Uint sets an uint value from an environment variable.
func Uint(p *uint, key string) error {
	return setValue((*uintValue)(p), key)
}

// Uint64 sets an uint64 value from an environment variable.
func Uint64(p *uint64, key string) error {
	return setValue((*uint64Value)(p), key)
}

// Float64 sets an float64 value from an environment variable.
func Float64(p *float64, key string) error {
	return setValue((*float64Value)(p), key)
}

// Text sets an [encoding.TextUnmarshaler] value from an environment variable.
func Text(p encoding.TextUnmarshaler, key string) error {
	return setValue(textValue{p}, key)
}

// Arg sets a [flag.Value] from an environment variable.
func Arg(value flag.Value, key string) error {
	return setValue(value, key)
}

// Enum sets a [flag.Value] from an environment variable,
// from a limited set of allowed values.
func Enum(value flag.Value, key string, values ...string) error {
	return setValue(newEnumValue(value, values), key)
}
