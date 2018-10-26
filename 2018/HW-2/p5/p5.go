/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 25-10-2018
 * |
 * | File Name:     p5.go
 * +===============================================
 */

package main

import (
	"fmt"
	"math"
)

// PI represents PI number approximation
const PI = 3.14

func main() {
	var ab float64  // ab side of the trapezium
	var abc float64 // abc angle of the trapezium in degree (>90)
	var bc float64  // bc side of the trapezium
	var bcd float64 // bcd angle of the trapezium in degree (>90)

	fmt.Scanf("%f %f %f %f", &ab, &abc, &bc, &bcd)

	if abc <= 90 || bcd <= 90 {
		panic("abc and bcd angles of the trapezium must be greater than 90 in degree")
	}

	var cd = math.Cos((abc-90)*PI/180) * ab * (1 / math.Cos((bcd-90)*PI/180)) // cd side of the trapezium
	var ad = bc + math.Sin((abc-90)*PI/180)*ab + math.Sin((bcd-90)*PI/180)*cd // ad side of the trapezium

	fmt.Printf("Perimeter: %.02f\n", ab+bc+cd+ad)
	fmt.Printf("Area: %.02f\n", math.Cos((abc-90)*PI/180)*ab*bc+0.5*math.Sin((bcd-90)*PI/180)*cd*math.Cos((bcd-90)*PI/180)*cd+0.5*math.Sin((abc-90)*PI/180)*ab*math.Cos((abc-90)*PI/180)*ab)
}
