package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// TODO : next feature
// *  Add possibility to reload a directory or add new directory at the runtime

// Albums Type  contain photos
type Albums struct {
	lstPhotos      []*Photo
	directoryPaths []string
}

// NewAlbum , return Albums struct
func NewAlbum(directoryPaths []string) *Albums {
	return &Albums{
		directoryPaths: directoryPaths,
	}
}

// LoadPhotosInAlbums , iterate for each directory and load photos
func (a *Albums) LoadPhotosInAlbums() (int, error) {
	// Loop in each directory and load picture files
	// TODO: need improve regex
	libRegEx, err := regexp.Compile("^.+\\.(jpg)$")
	if err != nil {
		return 0, err
	}

	// stop if no directory was provided
	if len(a.directoryPaths) < 1 {
		return 0, fmt.Errorf("No directory path provided")
	}

	// loop to extract picture for all directory
	for _, v := range a.directoryPaths {
		err = filepath.Walk(v, func(path string, info os.FileInfo, errFile error) error {
			if errFile == nil && libRegEx.MatchString(info.Name()) {
				err = a.chargePicInAlbum(path)
				if err != nil {
					fmt.Printf("Error with file %s error extracting tags : %v", path, err)
				}
			}
			return nil
		})
		if err != nil {
			fmt.Printf("Error processing file : %v", err)
		}
	}
	return len(a.lstPhotos), nil
}

//chargePicInAlbum , create a photo struct and load picture.
func (a *Albums) chargePicInAlbum(filename string) error {
	// init new photo
	photo := NewPhoto(filename)
	err := photo.LoadPhotoTags()
	if err != nil {
		return err
	}

	// convert raw data to the struct
	err = photo.SetPhotoStruct()
	if err != nil {
		return err
	}

	// everythings looks good, let's include this picture in the album
	a.lstPhotos = append(a.lstPhotos, photo)
	return nil
}

//PrintAlbumInfo , Print picture information
func (a *Albums) PrintAlbumInfo() error {
	// stop if no directory was provided
	if len(a.directoryPaths) < 1 {
		return fmt.Errorf("No directory path provided")
	}

	for _, photo := range a.lstPhotos {
		photo.PrintMainInfo()
	}
	return nil
}
