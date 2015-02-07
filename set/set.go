package main

import (
	"fmt"
	"github.com/deckarep/golang-set"
	"math/rand"
	"time"
)

func main() {

	// Prepare random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Printf("%s: ", time.Now())
	fmt.Println("Create two random sets")
	fmt.Println()

	set1 := mapset.NewSet()
	set2 := mapset.NewSet()

	for i := 0; i < 6000000; i++ {
		set1.Add(r.Intn(50000000))
		set2.Add(r.Intn(50000000))
	}

	fmt.Printf("%s: ", time.Now())
	fmt.Println("Finish create two random sets")
	fmt.Println()

	fmt.Printf("Elements in set1: ")
	fmt.Println(set1.Cardinality())

	fmt.Printf("Elements in set2: ")
	fmt.Println(set2.Cardinality())

	fmt.Printf("\n%s: ", time.Now())
	fmt.Println("Start union two random sets")
	fmt.Println()

	set3 := mapset.NewSet()

	set3 = set1.Union(set2)

	fmt.Printf("%s: ", time.Now())
	fmt.Println("Finish union two random sets")
	fmt.Println()

	fmt.Printf("Elements in set3: ")
	fmt.Println(set3.Cardinality())
}
