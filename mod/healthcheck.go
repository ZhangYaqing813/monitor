package mod

import (
	"github.com/idoubi/goz"
	"time"
)

func HealthCheck(uri string) (s string) {
	s = "OK"
	// create a curl instance
	cli := goz.NewClient(goz.Options{
		Timeout: 0.5,
	})

	_, err := cli.Get(uri)
	//错误处理
	if err != nil {
		// 如果uri 请求超时，则进额外行三次请求，全部超时后返回服务名称
		for i := 0; i < 3; i++ {
			time.Sleep(5 * time.Second)
			_, err := cli.Get(uri)
			if err != nil {
				continue

			} else {
				break
			}

		}
		s = "TimeOut"
		return s
	}

	return s
}
