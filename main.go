package main

import (
	"neutronstar-gb/hardware/cpu"
)

func main() {
	thecpu := cpu.CPU{}
	thecpu.Init()
	thecpu.Run()
}
