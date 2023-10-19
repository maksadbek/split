package split

import "testing"

func TestEqual(t *testing.T) {
	a := struct {
		A string
		B int
		C bool
	}{
		A: "A",
		B: 1,
		C: true,
	}

	b := struct {
		A string
		B int
		C bool
	}{
		A: "A",
		B: 1,
		C: true,
	}

	if !Equal(a, b) {
		t.Error("expected true, got false")
	}
}

func TestNotEqual(t *testing.T) {
	a := struct {
		A string
		B int
		C bool
	}{
		A: "A",
		B: 1,
		C: true,
	}

	b := struct {
		A string
		B int
		C bool
	}{
		A: "c",
		B: 2,
		C: false,
	}

	if Equal(a, b) {
		t.Error("expected false, got true")
	}

	t.Log(Diff(a, b))
}
