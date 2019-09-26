package webutil

import (
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) ([]byte, error) {
	// https://mp.weixin.qq.com/s?__biz=MjM5OTcxMzE0MQ==&mid=2653372872&idx=1&sn=19f914b95f15dc5f17799b247106b42e&scene=21#wechat_redirect

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //  关闭请求 归还连接

	rawContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return rawContent, nil
}
