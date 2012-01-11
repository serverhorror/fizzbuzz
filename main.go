package main

import (
	"./fizzbuzz"
	"flag"
)

var (
	upto *int = flag.Int("fizzbuzz", 20, "How long to play fizzbuzz (integer)")
)

func main() {
	flag.Parse()
	fizzbuzz.FizzBuzz(upto)
}
