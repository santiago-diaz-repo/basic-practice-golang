package main

import (
	"fmt"
	"testing"
)

func BenchmarkPrime(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		Prime(997)
	}
}

func BenchmarkPrime100(b *testing.B) {
	benchmarkPrime(100,b)
}

func benchmarkPrime(i int,b *testing.B) {
	for n := 0; n < b.N; n++ {
		Prime(i)
	}
}

func BenchmarkPrimeParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			Prime(2)
		}
	})
}

func Test_Sum(t *testing.T){
	fmt.Println("testing ***** ")
}