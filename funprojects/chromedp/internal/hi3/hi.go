package hi3

import (
	"github.com/chromedp/chromedp/device"
	"io/ioutil"
)

func main() {
	// 创建新的cdp上下文
	ctx, cancel := cdp.NewContext(context.Background())
	defer cancel()
	// run
	var b []byte
	if err := cdp.Run(ctx,
		// 模拟 iPhone 7
		cdp.Emulate(device.IPhone7landscape),
		cdp.Navigate(`https://www.whatsmyua.info/`),
		cdp.CaptureScreenshot(&b),
	); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("iphone7_ua.png", b, 0644); err != nil {
		log.Fatal(err)
	}
}
