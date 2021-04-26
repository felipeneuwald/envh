// Package envh provides convenient helpers to easily retrieve environment variables.
package envh

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// String receives a s string and retrieves the value of the environment variable named by s.
// It returns the value of the environment variable as a string.
// If the variable is not present in the environment, it returns an empty string and an error.
func String(s string) (string, error) {
	v, ok := os.LookupEnv(s)
	if !ok {
		err := fmt.Errorf("environment variable %v is missing", s)
		err = errors.New(err.Error())
		return "", err
	}
	return v, nil
}

// MustString receives a s string, passes it to func String, and returns the value returned by String.
// It differs from String because it does not return an error.
func MustString(s string) string {
	v, _ := String(s)
	return v
}

// MustStringFatal receives a s string, passes it to func String, and returns the value returned by String.
// If String returns an error, it prints an error message and exits.
func MustStringFatal(s string) string {
	v, err := String(s)
	if err != nil {
		log.Fatal().Stack().Err(err).Send()
	}
	return v
}

// Bool receives a s string and retrieves the value of the environment variable named by s.
// It expects "true" or "false" as the value of the environment variable and it returns the value as a bool.
// If the variable is not present in the environment or if the value is not expected, it returns false
// and an error.
func Bool(s string) (bool, error) {
	v, ok := os.LookupEnv(s)
	if !ok {
		err := fmt.Errorf("environment variable %v is missing", s)
		err = errors.New(err.Error())
		return false, err
	}
	if v == "true" {
		return true, nil
	}
	if v == "false" {
		return false, nil
	}
	err := fmt.Errorf("environment variable %v is %q; want \"true\" or \"false\"", s, v)
	err = errors.New(err.Error())
	return false, err
}

// MustBool receives a s string, passes it to func Bool, and returns the value returned by Bool.
// It differs from Bool because it does not return an error.
func MustBool(s string) bool {
	v, _ := Bool(s)
	return v
}

// MustBoolFatal receives a s string, passes it to func Bool, and returns the value returned by Bool.
// If Bool returns an error, it prints an error message and exits.
func MustBoolFatal(s string) bool {
	v, err := Bool(s)
	if err != nil {
		log.Fatal().Stack().Err(err).Send()
	}
	return v
}

// Int receives a s string and retrieves the value of the environment variable named by s.
// It returns the value of the environment variable as an int.
// If the variable is not present in the environment or strconv.Atoi can't convert the value
// to a string, it returns 0 and an error.
func Int(s string) (int, error) {
	v, ok := os.LookupEnv(s)
	if !ok {
		err := fmt.Errorf("environment variable %v is missing", s)
		err = errors.New(err.Error())
		return 0, err
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		err = errors.New(err.Error())
		return 0, err
	}
	return i, nil
}

// MustInt receives a s string, passes it to func Int, and returns the value returned by Int.
// It differs from Int because it does not return an error.
func MustInt(s string) int {
	v, _ := Int(s)
	return v
}

// MustIntFatal receives a s string, passes it to func Int, and returns the value returned by Int.
// If Int returns an error, it prints an error message and exits.
func MustIntFatal(s string) int {
	v, err := Int(s)
	if err != nil {
		log.Fatal().Stack().Err(err).Send()
	}
	return v
}
