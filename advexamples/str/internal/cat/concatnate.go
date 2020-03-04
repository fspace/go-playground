package cat

import (
	"bytes"
	"fmt"
	"strings"
)

func StringPlus() string {
	var s string
	s += "昵称" + ":" + "飞雪无情" + "\n"
	s += "博客" + ":" + "http://www.flysnow.org/" + "\n"
	s += "微信公众号" + ":" + "flysnow_org"
	return s
}

func StringFmt() string {
	return fmt.Sprint("昵称", ":", "飞雪无情", "\n", "博客", ":", "http://www.flysnow.org/", "\n", "微信公众号", ":", "flysnow_org")
}

func StringJoin() string {
	s := []string{"昵称", ":", "飞雪无情", "\n", "博客", ":", "http://www.flysnow.org/", "\n", "微信公众号", ":", "flysnow_org"}
	return strings.Join(s, "")
}

func StringBuffer() string {
	var b bytes.Buffer
	b.WriteString("昵称")
	b.WriteString(":")
	b.WriteString("飞雪无情")
	b.WriteString("\n")
	b.WriteString("博客")
	b.WriteString(":")
	b.WriteString("http://www.flysnow.org/")
	b.WriteString("\n")
	b.WriteString("微信公众号")
	b.WriteString(":")
	b.WriteString("flysnow_org")
	return b.String()
}

func StringBuilder() string {
	var b strings.Builder
	b.WriteString("昵称")
	b.WriteString(":")
	b.WriteString("飞雪无情")
	b.WriteString("\n")
	b.WriteString("博客")
	b.WriteString(":")
	b.WriteString("http://www.flysnow.org/")
	b.WriteString("\n")
	b.WriteString("微信公众号")
	b.WriteString(":")
	b.WriteString("flysnow_org")
	return b.String()
}
