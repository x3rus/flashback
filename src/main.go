package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Arg required, path to an image file")
		os.Exit(1)
	}

	// get file path
	photoFile := os.Args[1]
	fmt.Println(photoFile)

	photo := NewPhoto(photoFile)

	err := photo.LoadPhotoTags()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Extract All metadata:")
	photo.PrintAllMetaData()

	err = photo.SetPhotoStruct()
	if err != nil {
		fmt.Println("An error append with the data extraction")
		fmt.Println(err)

	}
	fmt.Printf("===================\n\n")
	fmt.Printf("GPS information : %s \n", photo.gps)
	fmt.Printf("Label associated : %s \n", photo.labels)
	fmt.Printf("Date when the picture was taken: %s \n", photo.dateCreation)

}
