package main

import (
	"fmt"
	"os"

	"github.com/mrTomatolegit/DFG/internal/parsers"
	"github.com/mrTomatolegit/DFG/internal/presentation"
	"github.com/mrTomatolegit/DFG/internal/util"
	"github.com/mrTomatolegit/DFG/internal/writers"
)

func main() {
	parsers.DefineFlags() // Define the flags for the help message
	presentation.StartTimer()
	if len(os.Args) == 1 {
		// No size arguments were provided, print instructions
		presentation.PrintHelp()
		return
	}
	flags := parsers.ParseFlags() // Parse the flags from the command line
	byteCount := flags.ByteCount
	output := flags.OutFile
	fileContent := flags.FileContent

	// Write the file (method is determined in the function)
	writers.Write(fileContent, byteCount, output, flags.WriteMem)

	fileInfo, err := os.Lstat(output) // Get file info
	util.Check(err)

	fmt.Println("\u001b[32m\n"+fmt.Sprint(fileInfo.Size()), "bytes written to", output) // Green
	presentation.PrintTimeElapsed()
	fmt.Print("\u001b[0m") // Reset
}
