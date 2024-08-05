package ozon_task

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type TestData struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

func loadTest(path string) (map[string]*TestData, error) {
	tests := make(map[string]*TestData)

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(path)
		filename := filepath.Base(strings.TrimSuffix(path, ext))

		_, ok := tests[filename]

		if !ok {
			tests[filename] = &TestData{}
		}

		body, err := os.ReadFile(path)

		if err != nil {
			return err
		}

		bodyString := string(body)

		// TODO: Прочитать файл

		if ext == ".a" {
			tests[filename].Output = bodyString
		}

		if ext == "" {
			tests[filename].Input = bodyString
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return tests, nil
}

func processTest(dir string, input string) (string, error) {
	fileName, err := findExeInDir(dir)

	if err != nil {
		return "", err
	}

	output, err := runExe(fileName, input)

	if err != nil {
		return "", err
	}

	return output, nil
}

func diff(a, b string) string {
	var result string

	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			result += fmt.Sprintf("Difference at position %d: '%c' vs '%c'\n", i, a[i], b[i])
		}
	}

	if len(a) > len(b) {
		result += fmt.Sprintf("Extra characters in string a: '%s'\n", a[len(b):])
	} else if len(b) > len(a) {
		result += fmt.Sprintf("Extra characters in string b: '%s'\n", b[len(a):])
	}

	return result
}
