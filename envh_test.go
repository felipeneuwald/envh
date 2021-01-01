package envh_test

import (
	"os"
	"testing"

	"github.com/felipeneuwald/envh"
)

func TestBool(t *testing.T) {
	os.Setenv("TEST1", "true")
	want := true
	got, err := envh.Bool("TEST1")
	if err != nil {
		t.Errorf("TEST1: Wasn't expecting err: %v", err)
	}
	if got != want {
		t.Errorf("TEST1: Bool() = %v; want %v", got, want)
	}

	os.Setenv("TEST2", "false")
	want = false
	got, err = envh.Bool("TEST2")
	if err != nil {
		t.Errorf("TEST1: Wasn't expecting err: %v", err)
	}
	if got != want {
		t.Errorf("TEST2: Bool() = %v; want %v", got, want)
	}

	os.Setenv("TEST3", "TRUE")
	want = false
	got, err = envh.Bool("TEST3")
	if err == nil {
		t.Errorf("TEST3: Was expecting err")
	}
	if got != want {
		t.Errorf("TEST3: Bool() = %v; want %v", got, want)
	}

	os.Setenv("TEST4", "True")
	want = false
	got, err = envh.Bool("TEST4")
	if err == nil {
		t.Errorf("TEST4: Was expecting err")
	}
	if got != want {
		t.Errorf("TEST4: Bool() = %v; want %v", got, want)
	}

	os.Setenv("TEST5", "FALSE")
	want = false
	got, err = envh.Bool("TEST5")
	if err == nil {
		t.Errorf("TEST5: Was expecting err")
	}
	if got != want {
		t.Errorf("TEST5: Bool() = %v; want %v", got, want)
	}

	os.Setenv("TEST6", "False")
	want = false
	got, err = envh.Bool("TEST6")
	if err == nil {
		t.Errorf("TEST6: Was expecting err")
	}
	if got != want {
		t.Errorf("TEST6: Bool() = %v; want %v", got, want)
	}

	os.Setenv("TEST7", "")
	want = false
	got, err = envh.Bool("TEST7")
	if err == nil {
		t.Errorf("TEST7: Was expecting err")
	}
	if got != want {
		t.Errorf("TEST7: Bool() = %v; want %v", got, want)
	}

	want = false
	got, err = envh.Bool("TEST8")
	if err == nil {
		t.Errorf("TEST8: Was expecting err")
	}
	if got != want {
		t.Errorf("TEST8: Bool() = %v; want %v", got, want)
	}
}

func TestMustBool(t *testing.T) {
	os.Setenv("TEST1", "true")
	want := true
	got := envh.MustBool("TEST1")
	if got != want {
		t.Errorf("TEST1: MustBool() = %v; want %v", got, want)
	}

	os.Setenv("TEST2", "false")
	want = false
	got = envh.MustBool("TEST2")
	if got != want {
		t.Errorf("TEST2: MustBool() = %v; want %v", got, want)
	}

	os.Setenv("TEST3", "TRUE")
	want = false
	got = envh.MustBool("TEST3")
	if got != want {
		t.Errorf("TEST3: MustBool() = %v; want %v", got, want)
	}

	os.Setenv("TEST4", "True")
	want = false
	got = envh.MustBool("TEST4")
	if got != want {
		t.Errorf("TEST4: MustBool() = %v; want %v", got, want)
	}

	os.Setenv("TEST5", "FALSE")
	want = false
	got = envh.MustBool("TEST5")
	if got != want {
		t.Errorf("TEST5: MustBool() = %v; want %v", got, want)
	}

	os.Setenv("TEST6", "False")
	want = false
	got = envh.MustBool("TEST6")
	if got != want {
		t.Errorf("TEST6: MustBool() = %v; want %v", got, want)
	}

	os.Setenv("TEST7", "")
	want = false
	got = envh.MustBool("TEST7")
	if got != want {
		t.Errorf("TEST7: MustBool() = %v; want %v", got, want)
	}

	want = false
	got = envh.MustBool("TEST8")
	if got != want {
		t.Errorf("TEST8: MustBool() = %v; want %v", got, want)
	}
}
