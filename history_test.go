package mazegen_test

import (
	"testing"

	"github.com/miketmoore/mazegen"
)

func TestNewHistory(t *testing.T) {
	history := mazegen.NewHistory()

	if history.Length() != 0 {
		t.Error("test failed")
	}

	// prove that this does not throw an error when
	// the slice is empty
	last := history.GetLast()
	if last != nil {
		t.Error("test failed")
	}

	// prove that this does not throw an error when
	// the slice is empty
	history.DeleteLast()

	history.Push(mazegen.NewCoordinates(1, 2))

	if history.Length() != 1 {
		t.Error("test failed")
	}

	last = history.GetLast()
	if last == nil {
		t.Error("test failed")
		return
	}
	if last.X != 1 && last.Y != 2 {
		t.Error("test failed")
	}

	history.DeleteLast()
	if history.Length() != 0 {
		t.Error("test failed")
	}
}
