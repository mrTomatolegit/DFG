package writers

import (
	"fmt"
	"github.com/mrTomatolegit/DFG/internal/util"
)

func Write(content util.FileContent, byteCount int, outFile string, writeMem int) {
	if byteCount > writeMem {
		fmt.Print("\u001b[33m") // Yellow
		fmt.Println("WARNING: The requested file generation is larger than the maximum memory assigned.")
		fmt.Println("Make sure you have enough Memory and Disk Space.")
		fmt.Print("Using a slower method to avoid using a lot of memory.\n\n")
		fmt.Print("\u001b[0m") // Reset
		ComplexWrite(content, byteCount, outFile, writeMem)
	} else {
		SimpleWrite(content, byteCount, outFile)
	}
}
