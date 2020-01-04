// strutil包 存放关于字符串的工具函数
package strutil

// @see https://segmentfault.com/a/1190000021426712

// Reverse 将参数中的字符串反转后的字符串
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
