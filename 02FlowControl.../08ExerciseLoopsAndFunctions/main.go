package main

import (
	"fmt"
)

// ニュートン法でxの平方根を計算する
func Sqrt(x float64) float64 {
	var z float64 = 1.0
	for i:=0; i<10; i++ {
		z = z- ((z*z - x) / (2 * z))
		fmt.Println(z)
	}
	return z

}

// ニュートン法でxの平方根を計算する (thresholdでループを抜ける)
func Sqrt_threshold(x float64) float64 {
	var z float64 = 1
	const threshold float64 = 1e-4

	for i:=0; i<10; i++ {
		zPlus := z -( (z*z - x) / (2*z) )
		fmt.Printf("[%2d], %f\n", i, z)
		if  (zPlus*zPlus - x) < threshold {
			return zPlus
		}
		z = zPlus
	}
	return -1
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt_threshold(2))
}
