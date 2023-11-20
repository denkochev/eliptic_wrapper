package wrapper

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"
)

// ! Wrapper works on secp256k1 curve
type ElipticPoint struct {
	X *big.Int
	Y *big.Int
}

// check if point is on secp256k1 curve
func (point *ElipticPoint) IsOnCurveCheck() bool {
	return btcec.S256().IsOnCurve(point.X, point.Y)
}

// print point coordinates
func (point *ElipticPoint) PrintPoint() {
	fmt.Printf("X - %x\n", point.X)
	fmt.Printf("Y - %x\n", point.Y)
}

// function to get basic point for secp256k1 function
func GetBasicG() (point ElipticPoint) {
	var G = ElipticPoint{}
	G.X = btcec.S256().Gx
	G.Y = btcec.S256().Gy
	return G
}

// set new point
func SetNewPoint(x, y *big.Int) ElipticPoint {
	var point ElipticPoint
	point.X = x
	point.Y = y
	return point
}

// set new point from hex numbers
func SetNewPointFromHex(x, y string) ElipticPoint {
	var point ElipticPoint
	point.X, point.Y = fromHex(x), fromHex(y)
	return point
}

// func that generates random point on secp256k1 curve
func GetRandomECpoint() (ElipticPoint, error) {
	var point ElipticPoint

	privateKey, err := btcec.NewPrivateKey()
	if err != nil {
		return point, errors.New("error in generating ec point")

	}

	rangom_point := privateKey.PubKey()

	point.X = rangom_point.X()
	point.Y = rangom_point.Y()
	return point, nil
}

// point A + point B = Point C
func AddElipticPoints(a, b ElipticPoint) ElipticPoint {
	resultPoint := ElipticPoint{}

	resultPoint.X, resultPoint.Y = btcec.S256().Add(a.X, a.Y, b.X, b.Y)

	return resultPoint
}

// double EC point
func Double(point ElipticPoint) ElipticPoint {
	doubled := ElipticPoint{}
	doubled.X, doubled.Y = btcec.S256().Double(point.X, point.Y)
	return doubled
}

// scalar multiply k * P
func ScalarMult(point ElipticPoint, k *big.Int) ElipticPoint {
	result := ElipticPoint{}
	result.X, result.Y = btcec.S256().ScalarMult(point.X, point.Y, k.Bytes())
	return result
}

/*
helpers
*/

func fromHex(s string) *big.Int {
	if s == "" {
		return big.NewInt(0)
	}
	r, ok := new(big.Int).SetString(s, 16)
	if !ok {
		panic("invalid hex in source file: " + s)
	}
	return r
}

func GenerateRandomBigInt(maxBits int) (*big.Int, error) {
	// Генерація випадкового числа
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(1).Lsh(big.NewInt(1), uint(maxBits)))
	if err != nil {
		return nil, err
	}

	return randomNumber, nil
}
