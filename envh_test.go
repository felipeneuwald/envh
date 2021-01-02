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

var boolTest = []boolTestType{
	{true, "TEST1", "true", true, false},
	{true, "TEST2", "false", false, false},
	{true, "TEST3", "TRUE", false, true},
	{true, "TEST4", "True", false, true},
	{true, "TEST5", "FALSE", false, true},
	{true, "TEST6", "False", false, true},
	{true, "TEST7", "", false, true},
	{false, "TEST8", "", false, true},
	{true, "TEST9", " ", false, true},
	{true, "TEST10", "S0m3th1ng", false, true},
}

func TestBool(t *testing.T) {
	for _, v := range boolTest {
		if v.envDeclare {
			os.Setenv(v.envKey, v.envValue)
		}

		got, err := envh.Bool(v.envKey)

		if v.wantErr && err == nil {
			t.Errorf("%v: Was expecting err", v.envKey)
		}

		if !v.wantErr && err != nil {
			t.Errorf("%v: Wasn't expecting err: %v", v.envKey, err)
		}

		if got != v.want {
			t.Errorf("%v: Bool() = %v; want %v", v.envKey, got, v.want)
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
			t.Errorf("%v: MustBool() = %v; want %v", v.envKey, got, v.want)
		}
	}
}
