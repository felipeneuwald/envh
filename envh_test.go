package envh_test

import (
	"os"
	"testing"

	"github.com/felipeneuwald/envh"
)

type boolTestType struct {
	envDeclare bool
	envKey     string
	envValue   string
	want       bool
	wantErr    bool
}

type intTestType struct {
	envDeclare bool
	envKey     string
	envValue   string
	want       int
	wantErr    bool
}

var boolTest = []boolTestType{
	{true, "BOOL1", "true", true, false},
	{true, "BOOL2", "false", false, false},
	{true, "BOOL3", "TRUE", false, true},
	{true, "BOOL4", "True", false, true},
	{true, "BOOL5", "FALSE", false, true},
	{true, "BOOL6", "False", false, true},
	{true, "BOOL7", "", false, true},
	{false, "BOOL8", "", false, true},
	{true, "BOOL9", " ", false, true},
	{true, "BOOL10", "S0m3th1ng", false, true},
}

var intTest = []intTestType{
	{true, "INT1", "9", 9, false},
	{true, "INT2", "", 0, true},
	{true, "INT3", " ", 0, true},
	{false, "INT4", "", 0, true},
	{true, "INT5", "SomeString", 0, true},
	{true, "INT6", "-1", -1, false},
}

func TestBool(t *testing.T) {
	for _, v := range boolTest {
		if v.envDeclare {
			os.Setenv(v.envKey, v.envValue)
		}

		got, err := envh.Bool(v.envKey)

		if v.wantErr && err == nil {
			t.Errorf("Bool(%q): Was expecting err", v.envKey)
		}

		if !v.wantErr && err != nil {
			t.Errorf("Bool(%q): Wasn't expecting err: %v", v.envKey, err)
		}

		if got != v.want {
			t.Errorf("Bool(%q) = %v; want %v", v.envKey, got, v.want)
		}
	}
}

func TestMustBool(t *testing.T) {
	for _, v := range boolTest {
		if v.envDeclare {
			os.Setenv(v.envKey, v.envValue)
		}

		got := envh.MustBool(v.envKey)

		if got != v.want {
			t.Errorf("MustBool(%q) = %v; want %v", v.envKey, got, v.want)
		}
	}
}

func TestInt(t *testing.T) {
	for _, v := range intTest {
		if v.envDeclare {
			os.Setenv(v.envKey, v.envValue)
		}

		got, err := envh.Int(v.envKey)

		if v.wantErr && err == nil {
			t.Errorf("Int(%q): Was expecting err", v.envKey)
		}

		if !v.wantErr && err != nil {
			t.Errorf("Int(%q): Wasn't expecting err: %v", v.envKey, err)
		}

		if got != v.want {
			t.Errorf("Int(%q) = %v; want %v", v.envKey, got, v.want)
		}
	}
}

func TestMustInt(t *testing.T) {
	for _, v := range intTest {
		if v.envDeclare {
			os.Setenv(v.envKey, v.envValue)
		}

		got := envh.MustInt(v.envKey)

		if got != v.want {
			t.Errorf("MustInt(%q) = %v; want %v", v.envKey, got, v.want)
		}
	}
}
