package handler

import (
	"testing"
)

func TestDataHandler(t *testing.T) {
	var tests = []struct {
		name      string
		repo_name string
	}{
		{"mylex", "LuckyWinner-Flutter"},
		{"yasirtaher", "string-prepend"},
		{"Kale71kale", "Kaleli71"},
		{"undefined", "undefined"},
		{"null", "null"},
	}

	for _, test := range tests {
		result, _ := dataHandler(test.name, test.repo_name)

		if result != nil {
			if result.StarCount < 0 {
				t.Error("Expected result at least 0 ")
			}
		}
	}
}
