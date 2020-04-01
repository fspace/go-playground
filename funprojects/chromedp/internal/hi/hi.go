package hi

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"io/ioutil"
	"log"
	"time"

	cdp "github.com/chromedp/chromedp"
)

func main() {
	// 创建新的cdp上下文
	ctx, cancel := cdp.NewContext(context.Background())
	defer cancel()

	// 此处以360搜索首页为例
	urlstr := `https://www.so.com/`
	var buf []byte
	// 需要截图的元素，支持CSS selector以及XPath query
	selector := `#main`
	if err := cdp.Run(ctx, elementScreenshot(urlstr, selector, &buf)); err != nil {
		log.Fatal(err)
	}
	// 写入文件
	if err := ioutil.WriteFile("360_so.png", buf, 0644); err != nil {
		log.Fatal(err)
	}
}

// 截图方法
func elementScreenshot(urlstr, sel string, res *[]byte) cdp.Tasks {
	return cdp.Tasks{
		// 打开url指向的页面
		cdp.Navigate(urlstr),

		// 等待待截图的元素渲染完成
		cdp.WaitVisible(sel, cdp.ByID),
		// 也可以等待一定的时间
		//cdp.Sleep(time.Duration(3) * time.Second),

		// 执行截图
		cdp.Screenshot(sel, res, cdp.NodeVisible, cdp.ByID),
	}
}

// 导出指定元素为PDF
func elementPDFPrint(urlstr, sel string, res *[]byte) cdp.Tasks {
	var err error
	return cdp.Tasks{
		cdp.Navigate(urlstr),
		cdp.Sleep(time.Duration(5) * time.Second),
		cdp.ActionFunc(func(ctx context.Context) error {
			// 获取pdf数据
			*res, _, err = page.PrintToPDF().Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}
