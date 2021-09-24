package human

import "strconv"

// BytesCount converts a bytes count in a human readable format
// Using IEC Standard
// The value is approximated (one digit )
func BytesCount(i uint64) string {
	const unit uint64 = 1024
	if i < unit {
		return strconv.FormatUint(i, 10) + " B"
	}
	div, exp := unit, 0
	for n := i / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	var str []byte
	str = strconv.AppendFloat(str, float64(i)/float64(div), 'f', 1, 64)
	return string(append(str, ' ', "KMGTPE"[exp], 'i', 'B'))
}
