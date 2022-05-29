package writers

import (
	"fmt"
	"github.com/mrTomatolegit/DFG/pkg/util"
	"math"
	"os"
	"strings"
)

func ComplexWrite(s string, byteCount int, outFile string) {
	var baseDataToWrite string

	if byteCount > len(s) {
		baseDataToWrite = s
	} else {
		baseDataToWrite = s[:byteCount]
	}

	f, err := os.Create(outFile)
	util.Check(err)
	defer f.Close()

	f.WriteString(baseDataToWrite)
	if byteCount > len(s) {
		var splitInto = int(math.Floor(math.Sqrt(float64(byteCount-len(baseDataToWrite))) / 8))
		var toWritePerWrite = byteCount / splitInto
		var missingBytes = byteCount - len(baseDataToWrite) - splitInto*toWritePerWrite

		fmt.Println("Writing", toWritePerWrite, "bytes per write for", splitInto, "writes")

		var charToRepeat = s[len(s)-1:]
		for i := 0; i < splitInto; i++ {
			_, err := f.WriteString(strings.Repeat(charToRepeat, toWritePerWrite))
			util.Check(err)
			f.Sync()
		}

		_, err := f.WriteString(strings.Repeat(charToRepeat, missingBytes))
		util.Check(err)
		f.Sync()
	}
}
