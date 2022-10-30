//go:build !default || all || benchmark
// +build !default all benchmark

package main

import (
	"testing"
	"time"
)

// Validate Time use to load photos in Album
func TestTimeLoadPhotosInAlbums(t *testing.T) {
	AlbumsTests := []struct {
		name                   string
		albumSample            string
		loadXtimeAlbum         int
		expectedImagesReturned int
		timeUsedSec            float64
	}{
		{"Load album with 50 pics",
			"../data/pic-sample/dir1/",
			5,
			50,
			7.0,
		},
		{"Load album with 100 pics",
			"../data/pic-sample/dir1/",
			10,
			100,
			13.0,
		},
		{"Load album with 200 pics",
			"../data/pic-sample/dir1/",
			20,
			200,
			28.0,
		},
		{"Load album with 500 pics",
			"../data/pic-sample/dir1/",
			50,
			500,
			83.0,
		},
	}
	// Loop in the list of photo
	for _, tt := range AlbumsTests {

		albumPaths := []string{}

		// ICI ICI ICI
		for i := 0; i < tt.loadXtimeAlbum; i++ {
			albumPaths = append(albumPaths, tt.albumSample)
		}

		album := NewAlbum(albumPaths)

		start := time.Now()
		numPicsLoaded, err := album.LoadPhotosInAlbums()

		elapsed := time.Since(start)

		if time.Duration(elapsed.Seconds()) > time.Duration(tt.timeUsedSec) {
			t.Errorf("Time use to load pics for test : \" %v \"  exceed the number expected ; got %f | expected : %f", tt.name, elapsed.Seconds(), tt.timeUsedSec)
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
