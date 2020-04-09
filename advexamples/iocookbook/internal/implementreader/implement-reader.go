package implementreader

import (
	"io"
	"os"
)

type myReader struct {
	content  []byte // 我们将从此处读取数据
	position int    // 从内容中已经读取的字节索引
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func (r *myReader) Read(buf []byte) (int, error) {
	remainingBytes := len(r.content) - r.position
	n := min(remainingBytes, len(buf))
	if n == 0 {
		return 0, io.EOF
	}
	copy(buf[:n], r.content[r.position:r.position+n])
	r.position += n
	return n, nil
}

func Main() {
	reader := myReader{content: []byte("this is the stuff I'm reading")}

	io.Copy(os.Stdout, &reader)
}
