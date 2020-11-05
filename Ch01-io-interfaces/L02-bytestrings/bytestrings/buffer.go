package bytestrings

import (
	"bytes"
	"io"
	"io/ioutil"
)

// Buffer demonstrates some tricks for initializing bytes
// These buffers implement an io.Reader interface
func Buffer(rawString string) *bytes.Buffer {
	// we'll start with a string encoded into raw bytes
	rawBytes := []byte(rawString)

	// there are a number of ways to create buffer from the raw bytes or from the original string
	var b = new(bytes.Buffer)
	b.Write(rawBytes)

	// alt
	b = bytes.NewBuffer(rawBytes)

	// and avoiding the initial byte array altogether
	b = bytes.NewBufferString(rawString)

	return b
}


// To string is an example of taking an io.Reader and consuming
// it all, then returning a string
func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}