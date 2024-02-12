package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

func main() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println(robotgo.GetMousePos())
	}
}
