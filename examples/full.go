package main

import (
	"fmt"

	"github.com/brunogui0812/sysprofiler"
)

func main() {
	hardware, _ := sysprofiler.Hardware()
	mem, _ := sysprofiler.Memory()
	display, _ := sysprofiler.Displays()
	audio, _ := sysprofiler.Audio()
	os, _ := sysprofiler.OS()
	applications, _ := sysprofiler.Applications()
	updates, _ := sysprofiler.Updates()

	fmt.Printf("Hardware: %v", hardware)
	fmt.Println("")
	fmt.Printf("Memory: %v", mem)
	fmt.Println("")
	fmt.Printf("Display: %v", display)
	fmt.Println("")
	fmt.Printf("Audio: %v", audio)
	fmt.Println("")
	fmt.Printf("OS: %v", os)
	fmt.Println("")
	fmt.Printf("Applications: %v", applications)
	fmt.Println("")
	fmt.Printf("Updates: %v", updates)
	fmt.Println("")

}
