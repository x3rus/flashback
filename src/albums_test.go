package main

import (
	"testing"
)

// Validate Photo tags extraction Raw with a subset of files
func TestLoadPhotoInAlbum(t *testing.T) {
	AlbumsTests := []struct {
		albumPath            []string
		expectedImagesLoaded int
		withError            bool
	}{
		{[]string{"../data/pic-sample/dir1/", "../data/unknow"},
			10,
			true,
		},
		{[]string{"../data/pic-sample/dir1/", "../data/pic-sample/dir2/"},
			20,
			false,
		},
	}
	// Loop in the list of photo
	for _, tt := range AlbumsTests {
		album := NewAlbum(tt.albumPath)
		numPicLoaded, err := album.LoadPhotosInAlbums()

		// ICI
		if (err != nil) == tt.withError {
			t.Errorf("Error not expected, when loading picture in the Albums %s; got error : %v", tt.albumPath, err)
			continue
		}

		if numPicLoaded != tt.expectedImagesLoaded {
			t.Errorf("Number of picture do not match what is expected for : %s : got %d want %d", tt.albumPath, numPicLoaded, tt.expectedImagesLoaded)
		}
	}
}
