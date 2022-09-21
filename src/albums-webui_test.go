package main

import (
	"testing"
)

// 		supportedPage: []string{"home", "listYear", "listMonth", "listDay"},

// Validate Photo tags extraction Raw with a subset of files
func TestUrlParse(t *testing.T) {
	URL2Tests := []struct {
		URL                 string
		expectedWebPageType string
		withError           bool
	}{
		{"/",
			"home",
			false,
		},
		{"/2020",
			"listYear",
			false,
		},
		{"/2020/02",
			"listMonth",
			false,
		},
		{"/2020/18",
			"listMonth",
			true,
		},
		{"/someTest",
			"",
			true,
		},
		{"/someTest/02/03",
			"",
			true,
		},
	}
	// Loop in the list of photo
	for _, tt := range URL2Tests {

		flashbackWebUI := NewAlbumWebUI(nil, nil)
		// TODO improvement add validation on the map[string]string
		webPageType, _, err := flashbackWebUI.parseURL(tt.URL)

		// ICI
		if err != nil && !tt.withError {
			t.Errorf("Error not expected, when parsing URL %s; got error : %v", tt.URL, err)
		}

		if webPageType != tt.expectedWebPageType {
			t.Errorf("Web page type does not match what is expected for : %s : got %s want %s", tt.URL, webPageType, tt.expectedWebPageType)
		}
	}
}
