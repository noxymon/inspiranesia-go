package service

import "fmt"

type HomeServiceComponent struct {
	name string
}

func newHomeServiceComponent(name string) {
	fmt.Println("Construct " + name)
}
