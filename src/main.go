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

	fmt.Println(" now print info :")
	photo.PrintAllMetaData()

	//	fmt.Printf("GPS information : %s ", photo.GetPhotoTag("GPSPosition"))
	//	fmt.Printf("GPS information : %s ", photo.GetPhotoTag("Keywords"))

}
