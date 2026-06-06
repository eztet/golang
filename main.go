package main

import (
	"fmt"
	"golang_study/feature1"
	"golang_study/feature2"

	"golang_study/feature_postgres/connection"
)

func main() {
	fmt.Println("Hello Git!")
	feature1.Feature1()
	feature2.Feature2()

	connection.CheckConnection()

}
