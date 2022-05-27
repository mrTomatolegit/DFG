package main

import (
	"fmt"
	"github.com/mrTomatolegit/DFG/internal/parsers"
	"github.com/mrTomatolegit/DFG/internal/presentation"
	"github.com/mrTomatolegit/DFG/pkg/util"
	"github.com/mrTomatolegit/DFG/pkg/writers"
	"math"
	"os"
)

const MASTER_STRING = "awo"

func main() {
	if len(os.Args) == 1 {
		// No size argument was provided, print instructions
		presentation.PrintHelp()
		return
	}

	byteCount, output := parsers.ParseInput()

	if float64(byteCount) >= (math.Pow(1024, 3)) {
		// The file is bigger than a gigabyte, print a warning
		fmt.Println("WARNING: The requested file generation is quite large. This may take a longer time to generate. Make sure you have enough Disk Space.")
		fmt.Println("Using a slower method to avoid memory overutilization.")
		writers.ComplexWrite(MASTER_STRING, byteCount, output)
	} else {
		// The file shouldn't be a danger to memory
		writers.SimpleWrite(MASTER_STRING, byteCount, output)
	}

	fileInfo, err := os.Lstat(output)
	util.Check(err)
	fmt.Println(fileInfo.Size(), "bytes written to", output)
}
