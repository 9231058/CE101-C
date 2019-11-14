package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)
func main() {
	var x1s,x1e,x2s,x2e,x3s,x3e,y1s,y1e,y2s,y2e,y3s,y3e int
	fmt.Scanf("%d %d\n%d %d\n%d %d\n%d %d\n%d %d\n%d %d", &x1s,&x1e,&y1s,&y1e,&x2s,&x2e,&y2s,&y2e,&x3s,&x3e,&y3s,&y3e)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var x1,x2,x3,y1,y2,y3 int
	// generating a random number between 0 - [end - start] + start   generated random numbers with Intn are in [0,n)
	x1 = r.Intn(x1e-x1s+1)+x1s
	x2 = r.Intn(x2e-x2s+1)+x2s
	x3 = r.Intn(x3e-x3s+1)+x3s
	y1 = r.Intn(y1e-y1s+1)+y1s
	y2 = r.Intn(y2e-y2s+1)+y2s
	y3 = r.Intn(y3e-y3s+1)+y3s
	fmt.Println(perimeter(x1,x2,x3,y1,y2,y3))

	// for correcting submissions easier
	// saving all possible answers in a array :
	xa1 := make([]int, x1e-x1s+1)
	for i := range xa1{
		xa1[i] = i+x1s
	}
	xa2 := make([]int, x2e-x2s+1)
	for i := range xa2{
		xa2[i] = i+x2s
	}
	xa3 := make([]int, x3e-x3s+1)
	for i := range xa3{
		xa3[i] = i+x3s
	}
	ya1 := make([]int, y1e-y1s+1)
	for i := range ya1{
		ya1[i] = i+y1s
	}
	ya2 := make([]int, y2e-y2s+1)
	for i := range ya2{
		ya2[i] = i+y2s
	}
	ya3 := make([]int, y3e-y3s+1)
	for i := range ya3{
		ya1[i] = i+y3s
	}
	var res []float64
	for _, x1i := range xa1{
		for _, x2i := range xa2{
			for _, x3i := range xa3{
				for _, y1i := range ya1{
					for _, y2i := range ya2{
						for _, y3i := range ya3{
							res = append(res, perimeter(x1i,x2i,x3i,y1i,y2i,y3i))
						}
					}
				}
			}
		}
	}

	fmt.Println(res)
}

func perimeter(x1,x2,x3,y1,y2,y3 int) float64 {
//	         	(x1,y1)
//		         |\
//		        | \
//		       |  \
//		      |   \
//		     |    \
//		    |     \
//	(x2,y2)|______\(x3.y3)
	d1 := math.Sqrt(math.Pow(float64(y2)-float64(y1),2) + math.Pow(float64(x2)- float64(x1),2))
	d2 := math.Sqrt(math.Pow(float64(y3)-float64(y1),2) + math.Pow(float64(x3)- float64(x1),2))
	d3 := math.Sqrt(math.Pow(float64(y3)-float64(y2),2) + math.Pow(float64(x3)- float64(x2),2))
	return d1+d2+d3
}