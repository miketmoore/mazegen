package mazegen_test

import (
	"testing"

	"github.com/miketmoore/mazegen"
)

func TestNewCell(t *testing.T) {
	cell := mazegen.NewCell()

	if cell == nil {
		t.Errorf("response is nil which is unexpected")
		return
	}

	if cell.Visited != false {
		t.Errorf("cell.Visited is not true which is unexpected")
	}

	if len(cell.Walls) != 4 {
		t.Errorf("cell.Walls map is not the correct length")
	}

	for wall, value := range cell.Walls {
		if value != true {
			t.Errorf("cell.Walls[%s] is not true which is unexpected", wall)
		}
	}
}
