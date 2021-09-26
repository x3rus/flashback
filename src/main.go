package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Arg required, path to a directory where your pick are")
		os.Exit(1)
	}

	// get file path
	dirPictures := os.Args[1]

	album := NewAlbum([]string{dirPictures})

	fmt.Println("Load All pics:")
	numLoadedPic, err := album.LoadPhotosInAlbums()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("We have loaded : %d , in directory : %s", numLoadedPic, dirPictures)

	err = album.PrintAlbumInfo()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
