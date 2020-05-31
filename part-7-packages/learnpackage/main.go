package main

import (
	"fmt"
	"learnpackage/simpleinterest" //importing custom package
	"log"
	"os"
	"strconv"
)

var p, r, t float64 = 5000.0, 10.0, 1.0

/*
* init function to check if p, r and t are greater than zero
 */
func init() {
	println("Main package initialized")

	if p < 0 {
		log.Fatal("Principal is less than zero")
	}
	if r < 0 {
		log.Fatal("Rate of interest is less than zero")
	}
	if t < 0 {
		log.Fatal("Duration is less than zero")
	}
	log.Print(fmt.Sprintf(`PRINCIPAL: %f | RATE: %f | TIME_IN_YEARS: %f`, p, r, t))
}

/**
* getSimpleInterestValues used to get values if informed in env vars
**/
func getSimpleInterestValues() (pi, ri, ti float64, err error) {
	pi, err = strconv.ParseFloat(os.Getenv("PRINCIPAL"), 64)
	if err != nil {
		pi = p
	}
	ri, err = strconv.ParseFloat(os.Getenv("RATE"), 64)
	if err != nil {
		ri = r
	}
	ti, err = strconv.ParseFloat(os.Getenv("TIME_IN_YEARS"), 64)
	if err != nil {
		ti = t
	}

	return
}

func main() {
	fmt.Println("Simple interest calculation")
	// I know that is a dumb way to get values, but this is just to explain the 'init()' function. Ignore it
	p, r, t, err := getSimpleInterestValues()
	if err != nil {
		log.Print(fmt.Sprintf(`Error '%s'`, err))
	}

	si := simpleinterest.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)
}
