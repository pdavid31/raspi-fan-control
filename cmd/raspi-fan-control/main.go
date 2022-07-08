package main

import (
	"fmt"
	"log"

	"github.com/pdavid31/raspi-fan-control/internal/temp"

	"github.com/stianeikeland/go-rpio/v4"
)

var (
	pwmPin = rpio.Pin(10)
)

func init() {
	if err := rpio.Open(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	temp, err := temp.Get()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(temp)
}
