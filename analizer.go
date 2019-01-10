package main

import (
"fmt"
"os"
"log"
"go/build"
"bufio"
)

func main() {
	if len(os.Args) < 2  {
		log.Fatal("There is not file to analize")
	}
	var identifier = os.Args[1]

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}

	s := string(os.PathSeparator)
	pathPrograms := s+"src"+s+"github.com"+s+"FERBPP"+s+"interpreter"+s+"programs"+s

	file, err := os.Open(gopath+pathPrograms+identifier)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	} 
}

func parse 
	