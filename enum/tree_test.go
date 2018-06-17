package enum

import (
	"github.com/workfit/tester/assert"
	"testing"
)

func TestBasicTree(t *testing.T) {

	set := NewSet()

	/*

		A (1)
			AA (4)
		B (2)
			BA (5)
			BB (6)
				BBA (10)
				BBB (11)
			BC (7)
		C (3)
			CA (8)
			CB (9)

	*/

	values := map[int]string{
		1:  "A",
		2:  "B",
		3:  "C",
		4:  "AA",
		5:  "BA",
		6:  "BB",
		7:  "BC",
		8:  "CA",
		9:  "CB",
		10: "BBA",
		11: "BBB",
	}

	parents := map[int]int{
		1:  0,
		2:  0,
		3:  0,
		4:  1,
		5:  2,
		6:  2,
		7:  2,
		8:  3,
		9:  3,
		10: 6,
		11: 6,
	}

	tree, err := set.AddTree("test", values, parents)

	assert.For(t).ThatActual(err).IsNil()
	assert.For(t).ThatActual(tree).IsNotNil()

}
