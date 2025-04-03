package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"github.com/spf13/pflag"
	"io"
	"os"
	"strings"
)

var (
	Help         bool
	Recursive    bool
	IgnoreCase   bool
	ArchiveTypes []string
	SearchPath   string
	SearchString string
)
var (
	SupportedArchiveTypes = []string{
		"jar",
		//"war",
		"zip",
		//"ear",
		//"rar",
		//"tar",
		//"gz",
		//"tgz",
		//"bz2",
		//"tbz",
		//"xz",
		//"txz",
		//"7z",
		//"deb",
		//"rpm",
	}
	SupportedArchiveTypesStr = strings.Join(SupportedArchiveTypes, ", ")
)

func init() {
	pflag.BoolVarP(&Help, "help", "h", false, "Help")
	pflag.BoolVarP(&Recursive, "recursive", "R", false, "Recursive")
	pflag.BoolVarP(&IgnoreCase, "ignore-case", "I", false, "Ignore Case")
	pflag.StringSliceVarP(&ArchiveTypes, "type", "t", SupportedArchiveTypes,
		"Archive Types to Search, Default All, Supported Types: "+SupportedArchiveTypesStr)
	pflag.Parse()
	if Help {
		pflag.Usage()
		os.Exit(0)
	}
	if err := ValidateInputs(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// main Example: go run tool\cmd\searchcomp\main.go Manifest "C:\envs\maven\repository\org\apache\poi\ooxml-schemas\1.3\ooxml-schemas-1.3.jar"
func main() {
	input := SearchPath
	keyword := []byte(SearchString)
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
			if !IgnoreCase && bytes.Contains(bs, keyword) {
				fmt.Println(f.Name)
			} else if IgnoreCase && bytes.Contains(bytes.ToLower(bs), bytes.ToLower(keyword)) {
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

func ValidateInputs() error {
	args := pflag.Args()
	if len(args) != 2 {
		switch len(args) {
		case 0:
			return errors.New("Missing Search Pattern or Path")
		case 1:
			return errors.New("Missing Search Pattern and Path")
		default:
			return errors.New("Too Much Inputs")
		}
	}
	SearchString = args[0]
	if len(strings.TrimSpace(SearchString)) == 0 {
		return fmt.Errorf("Invalida Search Pattern: %s", SearchString)
	}
	SearchPath = args[1]
	if len(strings.TrimSpace(SearchPath)) == 0 {
		return fmt.Errorf("Invalida Search Path: %s", SearchPath)
	}
	return nil
}
