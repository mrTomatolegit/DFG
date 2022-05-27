package parsers

import (
	"github.com/mrTomatolegit/DFG/pkg/util"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ParseInput() (int, string) {
	if !strings.Contains(strings.ToUpper(os.Args[1]), "B") {
		// No unit argument was provided
		os.Args[1] += "BY"
	}

	byteCount64, err := strconv.ParseInt(string(os.Args[1][:len(os.Args[1])-2]), 10, 32) // Get the file size provided
	util.Check(err)
	byteCount := int(byteCount64)
	measure := strings.ToUpper(os.Args[1][len(os.Args[1])-2:]) // Get the unit (provided or autofilled)

	// Get the output file (provided or autofilled):
	wd, err := os.Getwd()
	util.Check(err)

	var output string

	if len(os.Args) > 2 { // If file specified
		output = os.Args[2]
		if !strings.HasPrefix(output, "/") && !strings.HasPrefix(output, "./") {
			output = filepath.Join(wd, output)
		}
	} else { // Autofill
		output = filepath.Join(wd, "/output.dfg")
	}

	util.ApplyMeasurement(&byteCount, &measure)

	return byteCount, output
}
