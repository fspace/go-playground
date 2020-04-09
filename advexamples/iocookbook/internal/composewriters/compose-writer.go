package composewriters

import (
	"bytes"
	"io"
	"os"
)

type augmentedWriter struct {
	innerWriter io.Writer
	augmentFunc func([]byte) []byte
}

// replaces ' ' with '!'
func bangify(buf []byte) []byte {
	return bytes.Replace(buf, []byte(" "), []byte("!"), -1)
}

func (w *augmentedWriter) Write(buf []byte) (int, error) {
	return w.innerWriter.Write(w.augmentFunc(buf))
}

func BangWriter(w io.Writer) io.Writer {
	return &augmentedWriter{innerWriter: w, augmentFunc: bangify}
}

func UpcaseWriter(w io.Writer) io.Writer {
	return &augmentedWriter{innerWriter: w, augmentFunc: bytes.ToUpper}
}

func EncryptWriter(w io.Writer) io.Writer {
	// 从composereaders 包里面拷贝过来的 可以考虑共享
	encrypt := func(s []byte) []byte {
		result := make([]byte, len(s))

		for i, c := range s {
			result[i] = c + 28 //state-of-the-art encryption ladies and gentlemen
		}

		return result
	}

	return &augmentedWriter{
		innerWriter: w,
		augmentFunc: encrypt,
	}
}

func Main() {
	augmentedWriter := UpcaseWriter(BangWriter(os.Stdout))

	augmentedWriter.Write([]byte("lets see if this works\n")) // LETS!SEE!IF!THIS!WORKS
}
