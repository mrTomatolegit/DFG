package util

import (
	"math"
)

func ApplyMeasurement(byteCount *int, measure string) *int {
	switch measure {
	case "KB": // Kilobyte
		*byteCount *= int(math.Pow(1024, 1))
	case "MB": // Megabyte
		*byteCount *= int(math.Pow(1024, 2))
	case "GB": // Gigabyte
		*byteCount *= int(math.Pow(1024, 3))
	case "TB": // Terabyte
		*byteCount *= int(math.Pow(1024, 4))
	case "PB": // Petabyte (will use 1GB of memory)
		*byteCount *= int(math.Pow(1024, 5))
	}
	return byteCount
}
