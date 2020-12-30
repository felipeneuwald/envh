// Package envh is a helper for assigning environment variables.
package envh

import (
	"fmt"
	"os"
	"strconv"
)

// String receives a s string and retrieves the value of the environment variable named by s.
// It returns the value of the environment variable as a string.
// If the variable is not present in the environment, it returns an empty string and an error.
func String(s string) (string, error) {
	v, ok := os.LookupEnv(s)
	if !ok {
		return "", fmt.Errorf("Environment variable %v is missing", s)
	}
	return v, nil
}

// Bool receives a s string and retrieves the value of the environment variable named by s.
// It expects "true" or "false" as the value of the environment variable and it returns the value as a bool.
// If the variable is not present in the environment or if the value is not expected, it returns false
// and an error.
func Bool(s string) (bool, error) {
	v, ok := os.LookupEnv(s)
	if !ok {
		return false, fmt.Errorf("Environment variable %v is missing", s)
	}
	if v == "true" {
		return true, nil
	}
	if v == "false" {
		return false, nil
	}
	return false, fmt.Errorf("Environment variable %v is %q; want \"true\" or \"false\"", s, v)
}

// Int receives a s string and retrieves the value of the environment variable named by s.
// It returns the value of the environment variable as an int.
// If the variable is not present in the environment or strconv.Atoi can't convert the value
// to a string, it returns 0 and an error.
func Int(s string) (int, error) {
	v, ok := os.LookupEnv(s)
	if !ok {
		return 0, fmt.Errorf("Environment variable %v is missing", s)
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("Environment variable %v; %v", s, err)
	}
	return i, nil
}
