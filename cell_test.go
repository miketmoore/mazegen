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

func TestCarveWall(t *testing.T) {
	cell := mazegen.NewCell()
	cell.CarveWall(mazegen.North)

	if cell.Visited != true {
		t.Error("cell.Visited should be true but is not")
	}

	if cell.Walls[mazegen.North] != false {
		t.Error("cell.Walls[North] should be false but is not")
	}
}

func TestIsWallSolid(t *testing.T) {
	cell := mazegen.NewCell()
	got := cell.IsWallSolid(mazegen.North)

	if got != true {
		t.Error("response is false but should be true")
	}

	cell.CarveWall(mazegen.North)
	got = cell.IsWallSolid(mazegen.North)

	if got != false {
		t.Error("response is true but should be false")
	}
}

func TestOppositeDirection(t *testing.T) {
	cell := mazegen.NewCell()

	if cell.OppositeDirection(mazegen.North) != mazegen.South {
		t.Error("opposite of north is unexpected")
	}
	if cell.OppositeDirection(mazegen.East) != mazegen.West {
		t.Error("opposite of east is unexpected")
	}
	if cell.OppositeDirection(mazegen.West) != mazegen.East {
		t.Error("opposite of west is unexpected")
	}
	if cell.OppositeDirection(mazegen.South) != mazegen.North {
		t.Error("opposite of south is unexpected")
	}
	if cell.OppositeDirection("invalid") != "invalidDirection" {
		t.Error("invalid input has unexpected output")
	}
}
