package wrapper

import (
	"testing"
)

/*
test basic rules of eliptic curve Group

Idea:

k*(d*G) = d*(k*G)

ECPoint G = BasePointGGet()
big.Int k = SetRandom(256)
big.Int d = SetRandom(256)

H1 = ScalarMult(d, G)
H2 = ScalarMult(k, H1)

H3 = ScalarMult(k, G)
H4 = ScalarMult(d, H3)

bool result = IsEqual(H2, H4)

*/

func Test_Hypotesis(t *testing.T) {
	for i := 0; i < 10000; i++ {
		G := GetBasicG()

		k, _ := GenerateRandomBigInt(256)
		d, _ := GenerateRandomBigInt(256)

		H1 := ScalarMult(G, d)
		H2 := ScalarMult(H1, k)

		H3 := ScalarMult(G, k)
		H4 := ScalarMult(H3, d)

		result := H2.X.String() == H4.X.String() && H2.Y.String() == H4.Y.String()

		if !result {
			H2.PrintPoint()
			H4.PrintPoint()
			t.Error("Expected:- true, but got - false")
		}
	}
}
