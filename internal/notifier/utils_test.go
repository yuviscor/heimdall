package notifier

import (
	"fmt"
	"testing"
)

func Test_StatusCheck(t *testing.T) {

	TestCases := []int{200, 400, 300, 500}
	statusCodeTextMap := map[int]string{
		200: "Success",
		300: "Redirect",
		400: "Client Error",
		500: "Server Error",
	}

	for _, tc := range TestCases {

		t.Run("Status Code Text test", func(t *testing.T) {

			text := getStatusText(tc)

			if text != statusCodeTextMap[tc] {

				fmt.Println(statusCodeTextMap[tc])

				t.Errorf("Text to status code mapping is incorrect for %v", tc)
			}

		})

	}

}
