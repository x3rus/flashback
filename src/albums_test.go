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

// Validate when we query the album for a year you have your photos
func TestGetLstPhotosForYear(t *testing.T) {
	AlbumsTests := []struct {
		albumPath              []string
		expectedImagesReturned int
		yearToUse              int
	}{
		{[]string{"../data/pic-sample/dir2/", "../data/unknow"},
			1,
			2015,
		},
		{[]string{"../data/pic-sample/dir1/", "../data/pic-sample/dir2/"},
			7,
			2014,
		},
		{[]string{},
			0,
			2020,
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

		photoForTheYear, _ := album.GetLstPhotosForYear(tt.yearToUse)

		// for _, p := range photoForTheYear {
		//  	p.PrintMainInfo()
		//}

		if len(photoForTheYear) != tt.expectedImagesReturned {
			t.Errorf("Photo retreive for the date %d do not match number expected ; got %d | expected : %d", tt.yearToUse, len(photoForTheYear), tt.expectedImagesReturned)
		}

	}
}

// Validate Time use to load photos in Album
func TestTimeLoadPhotosInAlbums(t *testing.T) {
	AlbumsTests := []struct {
		albumPath              []string
		expectedImagesReturned int
		timeUsedSec            float64
	}{
		{[]string{"../data/pic-sample/dir2/", "../data/unknow"},
			6,
			2.0,
		},
		{[]string{"../data/pic-sample/dir1/", "../data/pic-sample/dir2/"},
			12,
			3.0,
		},
		// So if I can load multiple time the same directory I can simulate 100 dirs ...
		{[]string{"../data/pic-sample/dir1/", "../data/pic-sample/dir2/", "../data/pic-sample/dir2/", "../data/pic-sample/dir2/"},
			24,
			5.0,
		},
	}
	// Loop in the list of photo
	for _, tt := range AlbumsTests {
		album := NewAlbum(tt.albumPath)

		start := time.Now()
		numPicsLoaded, err := album.LoadPhotosInAlbums()

		elapsed := time.Since(start)

		if time.Duration(elapsed.Seconds()) > time.Duration(tt.timeUsedSec) {
			t.Errorf("Time use to load pics exceed the number expected ; got %f | expected : %f", elapsed.Seconds(), tt.timeUsedSec)
		}

		// for now I do not care the error validation performed in the previous test
		if err != nil {
			continue
		}

		if numPicsLoaded != tt.expectedImagesReturned {
			t.Errorf("Number of pics retreived do not match number expected ; got %d | expected : %d", numPicsLoaded, tt.expectedImagesReturned)
		}

	}

}

// inspiration from : https://blog.logrocket.com/benchmarking-golang-improve-function-performance/
// IMPROVEMENT: load different images
func BenchmarkChargePicInAlbum(b *testing.B) {
	album := NewAlbum([]string{"../data/pic-sample/dir1"})
	for i := 0; i < b.N; i++ {
		album.chargePicInAlbum("../data/pic-sample/dir1/20140410_150646.jpg")
	}
}
