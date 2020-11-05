package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

//ByFilename used for files sorting by name
type ByFilename []os.FileInfo

func (nf ByFilename) Len() int      { return len(nf) }
func (nf ByFilename) Swap(i, j int) { nf[i], nf[j] = nf[j], nf[i] }
func (nf ByFilename) Less(i, j int) bool {

	// Use path names
	a := strings.ToUpper(nf[i].Name())
	b := strings.ToUpper(nf[j].Name())

	// Which name is lexicographically smaller?
	return a < b
}

func main() {

	sorted := flag.Bool("s", false, "sort the elements")
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
	if *sorted {
		sort.Sort(ByFilename(files))
	}
	if *listed {
		for _, file := range files {
			fmt.Printf("%v %10v %9v%3v %v ", file.Mode(), file.Size(), file.ModTime().Month(),file.ModTime().Day(),file.ModTime().Format("15:04:05"))
			extension := filepath.Ext(path + "\\" + file.Name())
			if extension == "" {
				fmt.Println(string(colorBlue)+ file.Name()+string(colorReset)+"\\")
			} else if extension == ".dll" || extension == ".exe" {
				fmt.Println(string(colorGreen) + file.Name() + string(colorReset) + "*")
			} else {
				fmt.Println(file.Name())
			}
			fmt.Print(string(colorReset))
		}
	} else {
		for _, file := range files {
			extension := filepath.Ext(path + "\\" + file.Name())
			if extension == "" {
				fmt.Print(string(colorBlue), file.Name()+string(colorReset), "\\   ")
			} else if extension == ".dll" || extension == ".exe" {
				fmt.Print(string(colorGreen), file.Name()+string(colorReset), "*   ")
			} else {
				fmt.Print(file.Name() + "   ")
			}
			fmt.Print(string(colorReset))
		}
	}
}
