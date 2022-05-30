package parsers

import (
	"os"
	// "path/filepath"
	"strconv"
	"strings"

	"github.com/mrTomatolegit/DFG/internal/util"
)

func parseByteCount(s string) int {
	s = strings.ToUpper(s)
	if !strings.Contains(s, "B") {
		// No unit argument was provided
		s += "BY"
	}

	byteCount64, err := strconv.ParseInt(string(s[:len(s)-2]), 10, 32) // Get the file size provided
	util.Check(err)
	byteCount := int(byteCount64)
	measure := strings.ToUpper(s[len(s)-2:]) // Get the unit (provided or autofilled)

	util.ApplyMeasurement(&byteCount, measure)

	return byteCount
}

// Returns non ascii runes as '.'
func asciiOnly(r rune) rune {
	if int(r) > 255 { // Ascii characters end at 255
		return '.'
	}
	return r
}
func parseFileContent(content util.FileContent) util.FileContent {
	// Only allow ASCII characters
	// Unicode characters use more bytes than ASCII characters
	// which breaks the file size calculation
	content.Prefix = strings.Map(asciiOnly, content.Prefix)
	content.Repeat = strings.Map(asciiOnly, content.Repeat)
	content.Suffix = strings.Map(asciiOnly, content.Suffix)
	return content
}

type Flags struct {
	ByteCount   int
	OutFile     string
	FileContent util.FileContent
	WriteMem    int
}

// Gets the filesize which can be without flags as the first argument
func ParseBytesFlagless(s *string) {
	if !strings.HasPrefix(os.Args[1], "-") {
		*s = os.Args[1]
		os.Args = append(os.Args[1:], os.Args[2:]...) // Fix for flags only working on index 1
	}
}
