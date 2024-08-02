package byteutil

import (
	"strconv"
	"strings"
)

const (
	Byte = 1 << (10 * iota)
	Kilobyte
	Megabyte
	Gigabyte
	Terabyte
	Petabyte
	Exabyte
)

// ByteSize returns a human-readable byte string of the form 10M, 12.5K, and so forth.  The following units are available:
//
//	E: Exabyte
//	P: Petabyte
//	T: Terabyte
//	G: Gigabyte
//	M: Megabyte
//	K: Kilobyte
//	B: Byte
//
// The unit that results in the smallest number greater than or equal to 1 is always chosen.
func ByteSize(bytes uint64) string {
	unit := ""
	value := float64(bytes)

	switch {
	case bytes >= Exabyte:
		unit = "EiB"
		value = value / Exabyte
	case bytes >= Petabyte:
		unit = "PiB"
		value = value / Petabyte
	case bytes >= Terabyte:
		unit = "TiB"
		value = value / Terabyte
	case bytes >= Gigabyte:
		unit = "GiB"
		value = value / Gigabyte
	case bytes >= Megabyte:
		unit = "MiB"
		value = value / Megabyte
	case bytes >= Kilobyte:
		unit = "KiB"
		value = value / Kilobyte
	case bytes >= Byte:
		unit = "B"
	case bytes == 0:
		return "0 B"
	}

	// 保留小数点后四位
	result := strconv.FormatFloat(value, 'f', 4, 64)
	result = strings.TrimSuffix(result, ".0")
	return result + " " + unit
}
