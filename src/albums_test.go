package main

import (
	"testing"
	"time"
)

// Validate Photo tags extraction Raw with a subset of files
func TestLoadPhotoInAlbum(t *testing.T) {
	AlbumsTests := []struct {
		albumPath            []string
		expectedImagesLoaded int
		withError            bool
	}{
		{[]string{"../data/pic-sample/dir1/", "../data/unknow"},
			6,
			true,
		},
		{[]string{"../data/pic-sample/dir1/", "../data/pic-sample/dir2/"},
			12,
			false,
		},
		{[]string{},
			0,
			true,
		},
	}
	// Loop in the list of photo
	for _, tt := range AlbumsTests {
		album := NewAlbum(tt.albumPath)
		numPicLoaded, err := album.LoadPhotosInAlbums()

		// ICI
		if err != nil && !tt.withError {
			t.Errorf("Error not expected, when loading picture in the Albums %s; got error : %v", tt.albumPath, err)
		}

		if numPicLoaded != tt.expectedImagesLoaded {
			t.Errorf("Number of picture do not match what is expected for : %s : got %d want %d", tt.albumPath, numPicLoaded, tt.expectedImagesLoaded)
		}
	}
}

// Validate when we query the album with a date we have our photo
func TestGetLstPhotosForWeek(t *testing.T) {
	AlbumsTests := []struct {
		albumPath              []string
		expectedImagesReturned int
		dateToUse              time.Time
	}{
		{[]string{"../data/pic-sample/dir1/", "../data/unknow"},
			1,
			time.Date(2021, time.Month(12), 28, 1, 10, 30, 0, time.UTC),
		},
		{[]string{"../data/pic-sample/dir1/", "../data/pic-sample/dir2/"},
			2,
			time.Date(2020, time.Month(9), 15, 1, 10, 30, 0, time.UTC),
		},
		{[]string{},
			0,
			time.Date(2020, time.Month(4), 12, 1, 10, 30, 0, time.UTC),
		},
	}
	// Loop in the list of photo
	for _, tt := range AlbumsTests {
		album := NewAlbum(tt.albumPath)
		_, err := album.LoadPhotosInAlbums()

		// for now I do not care the error validation performed in the previous test
		if err != nil {
			continue
		}

		photoForTheDate, _ := album.GetLstPhotosForWeek(tt.dateToUse)

		//		for _, p := range photoForTheDate {
		//			p.PrintMainInfo()
		//		}

		if len(photoForTheDate) != tt.expectedImagesReturned {
			t.Errorf("Photo retreive for the data %s do not match number expected ; got %d | expected : %d", tt.dateToUse.String(), len(photoForTheDate), tt.expectedImagesReturned)
		}

	}
}
