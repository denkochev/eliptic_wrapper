package main

import (
	"fmt"
	"wrapper/wrapper"
)

func main() {
	// get G point of secp256k1 curve and print it
	G_point := wrapper.GetBasicG()
	G_point.PrintPoint()
	// check if point is on secp256k1 curve [always true]
	fmt.Println("Is G on point -> ", G_point.IsOnCurveCheck())

	// create random point and check if it's on curve [should be false with very high percentile]
	x, _ := wrapper.GenerateRandomBigInt(256)
	y, _ := wrapper.GenerateRandomBigInt(256)
	rand_point := wrapper.SetNewPoint(x, y)
	fmt.Println("Is random point belongs to secp256k1 curve ? - > ", rand_point.IsOnCurveCheck())

	// get random point ON secp256k1
	EC_point, err := wrapper.GetRandomECpoint()
	if err != nil {
		fmt.Println(err)
	}
	// should always be true
	fmt.Println("Is EC_point belongs to secp256k1 curve ? - > ", EC_point.IsOnCurveCheck())

	// create two points and do Group operations with them
	A_point, _ := wrapper.GetRandomECpoint()
	B_point, _ := wrapper.GetRandomECpoint()
	// add two point
	AB_point := wrapper.AddElipticPoints(A_point, B_point)
	AB_point.PrintPoint()

	// get double from point
	AB_double := wrapper.Double(AB_point)
	AB_double.PrintPoint()
	fmt.Println("Is doubled A+B point belongs to the curve? -> ", AB_double.IsOnCurveCheck())

	// multiply AB_double point on scalar
	scalar, _ := wrapper.GenerateRandomBigInt(256)

	AB_double_mul_scalar_point := wrapper.ScalarMult(AB_double, scalar)
	// check of AB_double_scalar_point belongs to the curve
	fmt.Println("Is (doubled AB) * scalar belongs to the curve -> ", AB_double_mul_scalar_point.IsOnCurveCheck())

}
