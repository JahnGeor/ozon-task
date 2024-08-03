package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

type TestData struct {
	output string
	input  string
}

func getTests(t *testing.T) map[string]TestData {
	testData := make(map[string]TestData)

	err := filepath.Walk("./test", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == "" {
			testData[info.Name()] = TestData{
				output: path + ".a",
				input:  path,
			}
		}

		return nil
	})

	if err != nil {
		t.Fatal(err)
	}

	return testData
}

func TestPath(t *testing.T) {
	tests := getTests(t)

	for name, test := range tests {
		fmt.Printf("Running test case №%s: %v \n", name, test)
		testTask2(test, t)
	}
}

func testTask2(testData TestData, t *testing.T) {

	cmd := exec.Command("go", "run", "main.go")

}
