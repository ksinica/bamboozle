package bamboozle_test

import (
	"testing"

	"github.com/ksinica/bamboozle"
)

func TestLipmaa(t *testing.T) {
	cases := map[uint64]uint64{
		1:   0,
		2:   1,
		3:   2,
		4:   1,
		5:   4,
		6:   5,
		7:   6,
		8:   4,
		9:   8,
		10:  9,
		11:  10,
		12:  8,
		13:  4,
		14:  13,
		15:  14,
		16:  15,
		17:  13,
		18:  17,
		19:  18,
		20:  19,
		21:  17,
		22:  21,
		23:  22,
		24:  23,
		25:  21,
		26:  13,
		27:  26,
		28:  27,
		29:  28,
		30:  26,
		31:  30,
		32:  31,
		33:  32,
		34:  30,
		35:  34,
		36:  35,
		37:  36,
		38:  34,
		39:  26,
		40:  13,
		121: 40,
	}
	for k, v := range cases {
		if bamboozle.Lipmaa(k) != v {
			t.Fail()
		}
	}
}
