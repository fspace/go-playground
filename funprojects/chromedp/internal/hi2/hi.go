package hi2

import (
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"io/ioutil"
	"math"
)

func Main() {
	// 创建新的cdp上下文
	ctx, cancel := cdp.NewContext(context.Background())
	defer cancel()

	// 此处以360官网首页为例
	urlstr := `https://www.360.cn/`
	var buf []byte
	// 获取 png, quality=90
	if err := cdp.Run(ctx, fullScreenshot(urlstr, 90, &buf)); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("360_cn_full.png", buf, 0644); err != nil {
		log.Fatal(err)
	}
}

func fullScreenshot(urlstr string, quality int64, res *[]byte) cdp.Tasks {
	return cdp.Tasks{
		cdp.Navigate(urlstr),
		cdp.ActionFunc(func(ctx context.Context) error {
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// 获取全屏截图
			*res, err = page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}
