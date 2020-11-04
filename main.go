package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	//"sort"
)

func main() {

	sorted := flag.Bool("S", false, "sort the elements")
	listed := flag.Bool("l", false, "list elements")
	flag.Parse()

	fmt.Println("sorted :", *sorted)
	fmt.Println("listed :", *listed)
	path := "C:\\Python38"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	colorBlue := "\033[34m"
	//sort.Strings(files)
	for _, file := range files {
		extension := filepath.Ext(path + "\\" + file.Name())
		if extension == "" {
			fmt.Print(string(colorBlue), file.Name(), string(colorReset), "\\   ")
		} else if extension == ".dll" || extension == ".exe" {
			fmt.Print(string(colorGreen), file.Name()+"   ")
		} else {
			fmt.Print(file.Name() + "   ")
		}
		//fmt.Println("Extension 1:", extension)
		fmt.Print(string(colorReset))
	}
}
