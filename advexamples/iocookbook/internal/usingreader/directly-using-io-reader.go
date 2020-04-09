package usingreader

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func Main() {
	reader := strings.NewReader("this is the stuff I'm reading")
	var result []byte
	buf := make([]byte, 4)
	for {
		n, err := reader.Read(buf)
		result = append(result, buf[:n]...)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}
	fmt.Println(string(result))
}
