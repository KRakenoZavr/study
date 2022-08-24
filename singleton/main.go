package main

import (
	"fmt"
)

type single struct {
	Count int
}

var singleInstance *single

func getInstance(i int) *single {
	if singleInstance == nil {
		fmt.Println("Creating single instance now")
		singleInstance = &single{Count: i}
	} else {
		fmt.Printf("Single instance already created: %v\n", singleInstance.Count)
	}

	return singleInstance
}

func main() {
	for i := 0; i < 20; i++ {
		getInstance(i)
	}
}
