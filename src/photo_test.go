package main

import (
	"fmt"
	"testing"
)

// Validate Photo tags extraction Raw with a subset of files
func TestTagsExtractionRaw(t *testing.T) {
	photoTests := []struct {
		filePath          string
		expectedTagsValue map[string]string
	}{
		{"../data/pic-sample/20131207_143248.jpg",
			map[string]string{
				"Keywords":    "Alice",
				"GPSPosition": "",
			},
		},
		{"../data/pic-sample/20160726_103655.jpg",
			map[string]string{
				"Keywords":    "[Alice Chalet vacance-2016 vacances]",
				"GPSPosition": "46 deg 5' 35.00\" N, 74 deg 48' 21.00\" W",
			},
		},
	}
	// Loop in the list of photo
	for _, tt := range photoTests {
		photo := NewPhoto(tt.filePath)
		err := photo.LoadPhotoTags()

		if err != nil {
			t.Errorf("Error Loading Tags for %s; got error : %v", tt.filePath, err)
			continue
		}

		// Tag extraction validation with direct access to the structure
		for tag, tagExpected := range tt.expectedTagsValue {
			if tagExpected == "" {
				if photo.metadata.Fields[tag] != nil {
					t.Errorf("For file : %s => Err extracting %s : got %s want %s", tt.filePath, tag, photo.metadata.Fields[tag], tagExpected)
				}
			} else if fmt.Sprintf("%v", photo.metadata.Fields[tag]) != tagExpected {
				t.Errorf("For file ; %s => Err extracting %s : got |%s| want |%s|", tt.filePath, tag, photo.metadata.Fields[tag], tagExpected)
			}
		}
	}
}

// Validate Photo tags extraction with a subset of files
func TestGetTags(t *testing.T) {
	photoTests := []struct {
		filePath          string
		expectedTagsValue map[string]string
	}{
		{"../data/pic-sample/20131207_143248.jpg",
			map[string]string{
				"Keywords":    "Alice",
				"GPSPosition": "",
			},
		},
		{"../data/pic-sample/20160726_103655.jpg",
			map[string]string{
				"Keywords":    "[Alice Chalet vacance-2016 vacances]",
				"GPSPosition": "46 deg 5' 35.00\" N, 74 deg 48' 21.00\" W",
			},
		},
	}
	// Loop in the list of photo
	for _, tt := range photoTests {
		photo := NewPhoto(tt.filePath)
		err := photo.LoadPhotoTags()

		if err != nil {
			t.Errorf("Error Loading Tags for %s; got error : %v", tt.filePath, err)
			continue
		}

		// Tag extraction validation with direct access to the structure
		for tag, tagExpected := range tt.expectedTagsValue {
			fmt.Printf(" process %s", tag)
			val := photo.GetPhotoTag(tag)
			if val != tagExpected {
				t.Errorf("For file ; %s => Err extracting %s : got |%s| want |%s|", tt.filePath, tag, photo.metadata.Fields[tag], tagExpected)
			}
		}
	}
}
