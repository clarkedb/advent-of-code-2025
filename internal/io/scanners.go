package io

import (
	"bytes"
)

// ScanCommas is a custom bufio.SplitFunc that reads text between commas
func ScanCommas(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// Look for a comma
	if i := bytes.IndexByte(data, ','); i >= 0 {
		// We found a comma at index i
		// return i+1 as the 'advance' (to move past the comma)
		// return data[0:i] as the token (the data up to the comma)
		return i + 1, data[0:i], nil
	}

	// If we are at EOF, we return the remaining data
	if atEOF {
		return len(data), data, nil
	}

	return 0, nil, nil
}
