package mod

import (
	"fmt"
	"github.com/idoubi/goz"
)

func HealthCheck() {

	service :=map[string]int{
		"goods":
	}

	cli := goz.NewClient(goz.Options{
		Timeout: 0.9,
	})
	resp, err := cli.Get("http://127.0.0.1:8091/get-timeout")
	if err != nil {
		if resp.IsTimeout() {
			fmt.Println("timeout")
			// Output: timeout
			return
		}
	}

}
