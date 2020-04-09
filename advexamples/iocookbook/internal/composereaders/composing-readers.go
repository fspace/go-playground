package composereaders

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

type augmentedReader struct {
	innerReader io.Reader
	augmentFunc func([]byte) []byte
}

// replaces ' ' with '!'
func bangify(buf []byte) []byte {
	return bytes.Replace(buf, []byte(" "), []byte("!"), -1)
}
func (r *augmentedReader) Read(buf []byte) (int, error) {
	tmpBuf := make([]byte, len(buf))
	n, err := r.innerReader.Read(tmpBuf)
	copy(buf[:n], r.augmentFunc(tmpBuf[:n]))
	return n, err
}

func BangReader(r io.Reader) io.Reader {
	return &augmentedReader{innerReader: r, augmentFunc: bangify}
}

func UpcaseReader(r io.Reader) io.Reader {
	return &augmentedReader{innerReader: r, augmentFunc: bytes.ToUpper}
}

func encrypt(s []byte) []byte {
	result := make([]byte, len(s))

	for i, c := range s {
		result[i] = c + 28 //state-of-the-art encryption ladies and gentlemen
	}

	return result
}

func EncryptReader(r io.Reader) io.Reader {
	return &augmentedReader{innerReader: r, augmentFunc: encrypt}
}

func Main() {
	originalReader := strings.NewReader("this is the stuff I'm reading")
	augmentedReader := UpcaseReader(BangReader(originalReader))

	result, err := ioutil.ReadAll(augmentedReader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(result)) // THIS!IS!THE!STUFF!I'M!READING
}
