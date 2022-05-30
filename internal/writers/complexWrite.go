package writers

import (
	"fmt"
	"github.com/mrTomatolegit/DFG/internal/presentation"
	"github.com/mrTomatolegit/DFG/internal/util"
	"os"
	"strings"
)

func writeString(f *os.File, s string, totalWritten *int) {
	b, err := f.WriteString(s)
	util.Check(err)
	*totalWritten += b
	f.Sync()
}

func ComplexWrite(content util.FileContent, byteCount int, outFile string, writeMem int) {
	var totalWritten int

	// Create a new file at the given path
	f, err := os.Create(outFile)
	util.Check(err)
	defer f.Close()

	var repStr = content.Repeat

	writer, quit := presentation.CreateUpdateTicker(&totalWritten, &byteCount)
	repStrCount := (byteCount - len(content.Prefix) - len(content.Suffix)) // How many bytes should each write use
	var splitInto = repStrCount / writeMem // How many writes to do
	var qToWrite = writeMem / len(repStr) // Amount for each write
	var remainder = (repStrCount % writeMem) / len(repStr) // Amount for the remaining write

	fmt.Println("Repeating", qToWrite*len(repStr), "bytes", splitInto, "times into", outFile)

	if len(content.Prefix) > 0 {
		// Add the prefix before we start adding the repStr
		writeString(f, content.Prefix, &totalWritten)
	}

	// Write the repStr the required amount of times
	for i := 0; i < splitInto; i++ {
		writeString(f, strings.Repeat(repStr, qToWrite), &totalWritten)
	}

	if remainder > 0 {
		// Add the remainder
		writeString(f, strings.Repeat(repStr, remainder), &totalWritten)
	}

	if totalWritten+len(content.Suffix) < byteCount {
		// Add the missing chars from part of rep (if rep length is odd)
		// We must take a single character to properly fix the missing bytes
		// If the repStr is of length 1 then this should never be called
		var lastChar = string(repStr[len(repStr)-1])
		var extraRequired = byteCount - totalWritten - len(content.Suffix) // Suffix will be added later so it needs to be accounted for
		fmt.Println("Adding", extraRequired, "extra bytes of", lastChar)
		writeString(f, strings.Repeat(lastChar, extraRequired), &totalWritten)
	}

	if len(content.Suffix) > 0 {
		// Add the suffix
		writeString(f, content.Suffix, &totalWritten)
	}

	presentation.UpdateProgressBar(writer, totalWritten, byteCount)
	close(quit)
}
