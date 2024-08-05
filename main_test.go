package ozon_task

import (
	"fmt"
	"testing"
)

func Test_loadTest(t *testing.T) {
	dir := "./task_4/"

	tests, err := loadTest(fmt.Sprintf("%s/test", dir))

	if err != nil {
		t.Fatal(err)
	}

	for key, test := range tests {
		if test.Output == "" || test.Input == "" {
			t.Fatalf("input/output %s is empty", key)
		}
	}

	for key, test := range tests {
		output, err := processTest(dir, test.Input)

		if err != nil {
			t.Fatalf("processTest #%s failed: %v", key, err)
		}

		difference := diff(output, test.Output)

		if diff(output, test.Output) != "" {
			t.Fatalf("processTest #%s failed: %s", key, difference)
		}
	}

}

func Test_processTest(t *testing.T) {

}
