package main

import (
	"fmt"
	"image"
	"image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"
)

func main() {
	var args string
	if len(os.Args) > 1 {
		args = os.Args[1]
	}

	openFile, err := os.Open(args)
	if err != nil {
		fmt.Println("Open file error.", err)
		os.Exit(0)
	}
	defer openFile.Close()

	img, _, err := image.Decode(openFile)
	if err != nil {
		fmt.Println("Read image error.", err)
		os.Exit(0)
	}

	fileSlice := strings.Split(args, ".")

	fileName := strings.Join(fileSlice[:len(fileSlice)-1], "") + ".gif"

	createFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Create new file error.", err)
		os.Exit(0)
	}
	defer createFile.Close()

	gif.Encode(createFile, img, &gif.Options{NumColors: 256})

	fmt.Printf("%s => %s  convert successful! \n", args, fileName)

}
