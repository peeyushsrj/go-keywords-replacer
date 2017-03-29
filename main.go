package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//https://gist.github.com/tdegrunt/045f6b3377f3f7ffa408
var count int

func visit(path string, fi os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !!fi.IsDir() {
		return nil //
	}
	matched, err := filepath.Match("*.php", fi.Name())
	if err != nil {
		log.Fatal(err)
		return err
	}
	if matched && path != "new.txt" && path != "old.txt" {
		read, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}
		newContents := string(read)
		for i, _ := range oldKwords {
			newContents = strings.Replace(newContents, oldKwords[i], newKwords[i], -1)
		}

		if newContents != string(read) {
			count++
			fmt.Println("Changed ", path)
		}

		err = ioutil.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

//https://gist.github.com/kendellfab/7417164
func readLine(path string) []string {
	var temp []string
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		temp = append(temp, scanner.Text())
	}
	fmt.Println("Imported keywords from ", path)
	return temp
}

var oldKwords []string
var newKwords []string

func main() {
	oldKwords = readLine("old.txt")
	newKwords = readLine("new.txt")
	err := filepath.Walk(".", visit)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count, " Matched files changed")
}
