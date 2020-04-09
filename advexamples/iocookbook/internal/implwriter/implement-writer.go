package implwriter

import "fmt"

type myWriter struct {
	content []byte
}

func (w *myWriter) Write(buf []byte) (int, error) {
	w.content = append(w.content, buf...)
	return len(buf), nil
}

func (w *myWriter) String() string {
	return string(w.content)
}

func Main() {
	writer := &myWriter{}
	writer.Write([]byte("hello there\n"))
	fmt.Println(writer.String())
}
