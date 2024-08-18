package main

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/fixme_my_friend/hw02_fix_app/employee"
	"github.com/fixme_my_friend/hw02_fix_app/printer"
	"github.com/fixme_my_friend/hw02_fix_app/reader"
)

const fileName = "data.json"

func main() {
	var err error
	var staff []employee.Employee
	staff, err = reader.ReadJSON(getFilePath(fileName))
	if err != nil {
		fmt.Println(err)
		return
	}
	printer.PrintStaff(staff)
}

func getFilePath(fileNameDefault string) string {
	var path string
	fmt.Printf("Enter data file path: ")
	_, err := fmt.Scanln(&path)
	if err != nil {
		return getPathByDefault(fileNameDefault)
	}
	if strings.TrimSpace(path) == "" {
		path = getPathByDefault(fileNameDefault)
	}
	return path
}

func getPathByDefault(fileName string) string {
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		fmt.Println("Error get the path to the source file")
	}

	directory := filepath.Dir(filename)
	fmt.Println(directory)
	return filepath.Join(directory, fileName)
}
