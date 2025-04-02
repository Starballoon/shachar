package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// main Example: go run tool\cmd\searchcomp\main.go "C:\envs\maven\repository\org\apache\poi\ooxml-schemas\1.3\ooxml-schemas-1.3.jar" Manifest
func main() {
	if len(os.Args) < 3 || len(strings.TrimSpace(os.Args[1])) == 0 || len(strings.TrimSpace(os.Args[2])) == 0 {
		fmt.Println("Invalid")
		return
	}
	input := os.Args[1]
	keyword := []byte(os.Args[2])
	if !IsFileExists(input) {
		fmt.Println("Invalid")
		return
	}
	r, err := zip.OpenReader(input)
	if err != nil {
		fmt.Println("Unable")
		return
	}
	defer r.Close()
	for _, f := range r.File {
		if !f.FileInfo().IsDir() {
			rc, err := f.Open()
			if err != nil {
				fmt.Println("Failed")
				return
			}
			bs, err := io.ReadAll(rc)
			if err != nil {
				fmt.Println("Failed")
				return
			}
			if bytes.Contains(bs, keyword) {
				fmt.Println(f.Name)
			}
			err = rc.Close()
			if err != nil {
				fmt.Println("Failed")
				return
			}
		}
	}
}

func IsFileExists(f string) bool {
	_, err := os.Stat(f)
	return err == nil
}
