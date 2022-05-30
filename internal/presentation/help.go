package presentation

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func PrintHelp() {
	var filename = strings.Split(filepath.Base(os.Args[0]), ".")[0]
	fmt.Printf("Usage: %s <size>[unit] \n", filename)
	fmt.Print("Example: dfg 8MB\n\n")
	fmt.Println("<size> is the size of the file to generate.")
	fmt.Println("[unit] is the unit of the file size. Valid units are: BY, KB, MB, GB, TB, PB. Defaults to \"BY\"")
	fmt.Println("\nFlags:")
	flag.PrintDefaults()
}
