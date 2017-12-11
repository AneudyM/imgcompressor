package main

import (
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"
	"log"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("You need to specify an image or directory.")
	}

	src := os.Args[1]

	filepath.Walk(src, compressJpeg)
}

func compressJpeg(path string, info os.FileInfo, err error) error {
	// Skip processing directory names including CWD
	if info.IsDir() {
		return nil
	}

	filename := info.Name()

	if isJpeg(filename) == false {
		return nil
	}

	quality := 70
	options := jpeg.Options{
		Quality: quality,
	}
	targetFile := fmt.Sprintf("%s", filename)

	sourcefile, err := os.Open(path)
	checkError(err)
	defer sourcefile.Close()

	img, err := jpeg.Decode(sourcefile)
	if err != nil {
		fmt.Println("File: " + info.Name())
		log.Fatal(err)
	}

	_, err = os.Stat("img")
	if os.IsNotExist(err) {
		err = os.MkdirAll("img", os.ModePerm)
		checkError(err)
	}

	target, err := os.Create(filepath.Join("img", targetFile))
	checkError(err)

	err = jpeg.Encode(target, img, &options)
	checkError(err)

	return nil
}

func isJpeg(file string) bool {
	fileExtension := filepath.Ext(file)

	if (fileExtension != ".jpg") && (fileExtension != ".jpeg") {
		return false
	} else {
		return true
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}