package testDemo

import "testing"

func TestAdd(t *testing.T) {
	got := Add(1, 2)
	want := 1 + 2

	if got != want {
		t.Errorf("expcted:%v, got:%v", want, got)
	}
}

func TestSquare(t *testing.T) {
	got := Square(5)
	want := 5 * 5

	if got != want {
		t.Errorf("expcted:%v, got:%v", want, got)
	}
}
