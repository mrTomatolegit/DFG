package parsers

import (
	"flag"

	"github.com/mrTomatolegit/DFG/internal/util"
)

var flagsDefined = false
var fa [6]*string // flag addresses

func DefineFlags() (*string, *string, *string, *string, *string, *string) {
	if !flagsDefined {
		flagsDefined = true
		fa[0] = flag.String("size", "1BY", "The size of the output file (BY, KB, MB, GB, TB, PB)")
		fa[1] = flag.String("out", "output.dfg", "The output file to write to")
		fa[2] = flag.String("pre", "aw", "The prefix to the file data")
		fa[3] = flag.String("rep", "o", "The string to repeat in the file data")
		fa[4] = flag.String("suf", "", "The suffix to the file data")
		fa[5] = flag.String("mem", "512MB", "The maximum amount of saving memory to use for each write")
	}

	return fa[0], fa[1], fa[2], fa[3], fa[4], fa[5]
}

func ParseFlags() Flags {
	bytes,
		out,
		prefix,
		repstr,
		suffix,
		writemem := DefineFlags()
	ParseBytesFlagless(bytes)
	flag.Parse()
	return Flags{
		ByteCount:   parseByteCount(*bytes),
		OutFile:     *out,
		FileContent: parseFileContent(util.FileContent{Prefix: *prefix, Repeat: *repstr, Suffix: *suffix}),
		WriteMem:    parseByteCount(*writemem),
	}
}
