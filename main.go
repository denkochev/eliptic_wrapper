package main

import (
	"fmt"
	"wrapper/wrapper"
)

func main() {
	// pring secp256k1 G coordinates
	wrapper.PrintBasicG()

	x, _ := wrapper.GenerateRandomBigInt(256)
	y, _ := wrapper.GenerateRandomBigInt(256)

	point := wrapper.SetNewPoint(x, y)
	// check if our random point on secp256k1
	fmt.Println("Is point on secp256k1 curve ? - > ", point.IsOnCurveCheck())

	// get random point on secp256k1
	EC, err := wrapper.GetRandomECpoint()
	if err != nil {
		fmt.Println(err)
	}
	// should always be true
	fmt.Println("Is EC on secp256k1 curve ? - > ", EC.IsOnCurveCheck())

	A_point, _ := wrapper.GetRandomECpoint()
	B_point, _ := wrapper.GetRandomECpoint()
	// add two point
	AB_point := wrapper.AddElipticPoints(A_point, B_point)
	AB_point.PrintPoint()

	// get double from point
	AB_double := wrapper.Double(AB_point)
	AB_double.PrintPoint()
	fmt.Println(AB_double.IsOnCurveCheck())
}
