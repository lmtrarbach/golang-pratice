package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	number_of_lines_per_file := 10000
	var file_number, count int = 1, 1
	fileToBeChunked := "./annual-enterprise-survey-2019-financial-year-provisional-csv.csv"
	fileList := make([]string, 0)
	f, _ := os.Open(fileToBeChunked)
	scanner := bufio.NewScanner(f)
	// Loop over lines in file.

	for scanner.Scan() {
		fmt.Printf("Count =  %v", count)
		fileList = append(fileList, scanner.Text())
		count += 1
		if count%number_of_lines_per_file == 0 {
			fileName := "somebigfile_" + strconv.Itoa(file_number)
			file_number += 1
			_, err := os.Create(fileName)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			// write/save buffer to disk
			ioutil.WriteFile(fileName, []byte(strings.Join(fileList, "\n")), os.ModeAppend)
			fmt.Println("Split to : ", fileName)
			fileList = nil
		}

	}

}

