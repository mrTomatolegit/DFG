package writers

import (
	"fmt"
	"os"
	"strings"
)

func SimpleWrite(s string, byteCount int, outFile string) {
	fmt.Println("Writing", byteCount, "bytes to", outFile)
	if byteCount > len(s) {
		repChar := string(s[len(s)-1])
		os.WriteFile(outFile, []byte(s+strings.Repeat(repChar, byteCount-len(s))), 0644)
	} else {
		os.WriteFile(outFile, []byte(s[:byteCount]), 0644)
	}
}
