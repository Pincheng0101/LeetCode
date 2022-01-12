package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	var problemTitle string
	flag.StringVar(&problemTitle, "title", "", "Title of the problem")
	flag.Parse()

	data := strings.Fields(problemTitle)
	v, err := strconv.Atoi(data[0][:len(data[0])-1])
	if err != nil {
		fmt.Println("data[0] convert integer failed: ", err.Error())
		os.Exit(0)
	}
	number := fmt.Sprintf("%04d", v)
	folderName := number + "." + strings.Join(data[1:], "-")
	folderPath := "./problems/" + folderName

	err = os.Mkdir(folderPath, 0755)
	if err != nil {
		fmt.Println("mkdir failed: ", err.Error())
		os.Exit(0)
	}
	fmt.Println("create folder: ", folderPath)
	file := fmt.Sprintf("%s/%s.go", folderPath, number)
	err = ioutil.WriteFile(file, []byte(fmt.Sprintf("package p%s\n", number)), 0755)
	if err != nil {
		fmt.Println("create file failed: ", err.Error())
		os.Exit(0)
	}
	fmt.Println("create file: ", file)
	testFile := fmt.Sprintf("%s/%s_test.go", folderPath, number)
	err = ioutil.WriteFile(testFile, []byte(fmt.Sprintf("package p%s\n", number)), 0755)
	if err != nil {
		fmt.Println("create file failed: ", err.Error())
		os.Exit(0)
	}
	fmt.Println("create test file: ", testFile)
	os.Exit(0)
}
