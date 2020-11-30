package main

import (
	"github.com/cgimenes/gomenes-boy/hardware/cpu"
)

func main() {
	thecpu := cpu.CPU{}
	thecpu.Init()
	thecpu.Run()
}
