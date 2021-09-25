package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
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
func TestTagsExtracted(t *testing.T) {
	photoTests := []struct {
		filePath            string
		labelsExpected      []string
		GPSPositionExpected string
		dateCreation        time.Time
	}{
		{"../data/pic-sample/20131207_143248.jpg",
			[]string{"Alice"},
			"",
			time.Date(2013, time.December, 07, 14, 32, 47, 0, time.UTC),
		},
		{"../data/pic-sample/20160726_103655.jpg",
			[]string{"Alice", "Chalet", "vacance-2016", "vacances"},
			"46 deg 5' 35.00\" N, 74 deg 48' 21.00\" W",
			time.Date(2016, time.July, 26, 10, 36, 55, 0, time.UTC),
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

		err = photo.SetPhotoStruct()
		if err != nil {
			t.Errorf("For file ; %s => Err extracting fields : got error |%s|", tt.filePath, err)
		}
		// Validate Label
		if !reflect.DeepEqual(photo.labels, tt.labelsExpected) {
			t.Errorf("For file ; %s => Err extracting Label: got |%s| want |%s|", tt.filePath, photo.labels, tt.labelsExpected)
		}

		// Validate GPGLocation
		if photo.gps != tt.GPSPositionExpected {
			t.Errorf("For file ; %s => Err extracting GPS Location: got |%s| want |%s|", tt.filePath, photo.gps, tt.GPSPositionExpected)
		}

		// Validate DateCreation
		if !tt.dateCreation.Equal(photo.dateCreation) {
			t.Errorf("For file ; %s => Err extracting date Creation: got |%s| want |%s|", tt.filePath, photo.dateCreation, tt.dateCreation)
		}
	}
}
