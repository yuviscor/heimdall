package domain

import (
	"testing"
)

func Test_ParseInterval(t *testing.T) {

	TestCases := []Service{
		{

			Name:     "TestIntervalString",
			Interval: "5m",
		},
		{

			Name:     "TestNumber",
			Interval: "1",
		},
		{

			Name:     "TestRandom",
			Interval: "1m",
		},
	}

	for _, tc := range TestCases {

		t.Run(tc.Name, func(t *testing.T) {

			_, err := tc.ParseInterval()

			if err != nil {

				t.Errorf("Test case failed: Failed to parse the Interval %v", err)
			}

		})

	}

}

func Test_ParseIntervalOnlyString(t *testing.T) {

	TestCases := []Service{
		{

			Name:     "TestIntervalString",
			Interval: "m",
		},
		{

			Name:     "TestNumber",
			Interval: "s",
		},
	}

	for _, tc := range TestCases {

		t.Run(tc.Name, func(t *testing.T) {

			_, err := tc.ParseInterval()

			if err == nil {

				t.Error("Only String Interval pasrse, Test Case Failed")
			}

		})

	}

}
