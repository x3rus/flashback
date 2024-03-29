package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

// TODO : next feature
// *  Add possibility to reload a directory or add new directory at the runtime

// WeeklyPics , slice to store pic together for the same month
type WeeklyPics map[int][]*Photo

// YearPics , slice to store pic together for the same year and month
type YearPics map[int]WeeklyPics

// Albums Type  contain photos
type Albums struct {
	lstPhotos      YearPics
	directoryPaths []string
	numPic         int
}

// NewAlbum , return Albums struct
func NewAlbum(directoryPaths []string) *Albums {
	return &Albums{
		directoryPaths: directoryPaths,
		lstPhotos:      make(YearPics),
		numPic:         0,
	}
}

// LoadPhotosInAlbums , iterate for each directory and load photos
// Return number of pictures loaded
func (a *Albums) LoadPhotosInAlbums() (int, error) {
	// Loop in each directory and load picture files
	// TODO: need improve regex
	libRegEx, err := regexp.Compile(`.+\.(jpg)$`)
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
	return a.numPic, nil
}

// chargePicInAlbum , create a photo struct and load picture.
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

	a.addPhoto(photo)
	return nil
}

// addPhoto , add the picture in the global album
// TODO : check how I can add error management
func (a *Albums) addPhoto(photo *Photo) {
	// extract date information
	year, week := photo.dateCreation.ISOWeek()

	// check if year was already created
	if _, ok := a.lstPhotos[year]; !ok {
		a.lstPhotos[year] = make(WeeklyPics)
	}
	// everythings looks good, let's include this picture in the album
	a.lstPhotos[year][int(week)] = append(a.lstPhotos[year][int(week)], photo)

	a.numPic++

}

// PrintAlbumInfo , Print picture information
func (a *Albums) PrintAlbumInfo(w io.Writer) error {
	// stop if no directory was provided
	if len(a.directoryPaths) < 1 {
		return fmt.Errorf("No directory path provided")
	}

	// Go throw the album and print information for each pictures
	for k := range a.lstPhotos {
		fmt.Printf(" We process year : %d \n", k)
		for m := range a.lstPhotos[k] {
			for _, photo := range a.lstPhotos[k][m] {
				photo.PrintMainInfo(w)
			}
		}
	}
	return nil
}

// GetLstPhotosForYear , Query the albums to retreive photo taken during the  year
func (a *Albums) GetLstPhotosForYear(year int) ([]*Photo, error) {

	// variable to hold all picture
	var photos []*Photo

	if photosOfTheYear, ok := a.lstPhotos[year]; ok {
		for p := range photosOfTheYear {
			photos = append(photos, photosOfTheYear[p]...)
		}
		return photos, nil
	}

	// if no pics found
	return nil, fmt.Errorf("No pictures found for this year ")
}

// GetLstPhotosForWeek , Query the albums to retreive photo taken during the week of dateSelected
func (a *Albums) GetLstPhotosForWeek(dateSelected time.Time) ([]*Photo, error) {
	// retrieve information from the data provided
	_, week := dateSelected.ISOWeek()

	// variable to hold all picture
	var photoOfTheWeeks []*Photo

	// Go throw the album and print information for each pictures
	// TODO : I think there is space here for optimization
	for k := range a.lstPhotos {
		for w := range a.lstPhotos[k] {
			if w == week {
				photoOfTheWeeks = append(photoOfTheWeeks, a.lstPhotos[k][w]...)
			}
		}
	}

	if len(photoOfTheWeeks) < 1 {
		// Not sure for the error it's not really an issue. TODO review the error
		return nil, fmt.Errorf("No pictures found for this year and week")
	}

	return photoOfTheWeeks, nil
}
