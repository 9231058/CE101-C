package main

import (
	"fmt"
	"math"
)

func main() {
	var xa, ya int
	var xb, yb int
	var xd, yd int

	fmt.Scanf("%d %d", &xa, &ya)
	fmt.Scanf("%d %d", &xb, &yb)
	fmt.Scanf("%d %d", &xd, &yd)

	ux := xb - xa
	uy := yb - ya

	vx := xd - xa
	vy := yd - ya

	cu := ux*ux + uy*uy
	cv := vx*vx + vy*vy
	ip := (ux*vx + uy*vy)

	op := cu*cv - ip*ip

	fmt.Printf("%.2f\n", math.Sqrt(float64(op)))
}
