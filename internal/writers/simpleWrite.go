package writers

import (
	"fmt"
	"os"
	"strings"

	"github.com/mrTomatolegit/DFG/internal/util"
)

func SimpleWrite(content util.FileContent, byteCount int, outFile string) {
	fmt.Println("Writing", byteCount, "bytes to", outFile)

	s := content.Prefix + content.Suffix

	if byteCount > len(s) {
		repStr := content.Repeat
		repStrCount := (byteCount - len(content.Prefix) - len(content.Suffix)) / len(repStr)
		writeString := content.Prefix + strings.Repeat(repStr, repStrCount)
		if len(writeString)+len(content.Suffix) < byteCount {
			// Add the missing chars from part of rep (if rep length is odd)
			writeString += repStr[:(byteCount - len(writeString) - len(content.Suffix))]
		}
		writeString += content.Suffix
		os.WriteFile(outFile, []byte(writeString), 0644)
	} else {
		os.WriteFile(outFile, []byte(s[:byteCount]), 0644)
	}
}
