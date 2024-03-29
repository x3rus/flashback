package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

// Page struct for the web server
type Page struct {
	Title string
	Body  []byte
}

// AlbumWebUI web interface for the UI
type AlbumWebUI struct {
	album         *Albums
	supportedPage []string
	appPrefix     string
	logger        *logrus.Entry
}

// NewAlbumWebUI , return Albums struct
func NewAlbumWebUI(album *Albums, appPrefix string, logger *logrus.Entry) *AlbumWebUI {
	return &AlbumWebUI{
		album:         album,
		supportedPage: []string{"home", "listYear", "listMonth", "listDay"},
		appPrefix:     appPrefix,
		logger:        logger,
	}
}

// RunWebSrv , provide the web interface
func (a *AlbumWebUI) RunWebSrv() {

	fs := http.FileServer(http.Dir("./data"))
	http.Handle("/data/", http.StripPrefix("/data/", fs))

	http.HandleFunc(a.appPrefix, a.viewHandler)
	a.logger.Fatal(http.ListenAndServe(":8080", nil))
}

// loadPage , show web page for the user
func (a *AlbumWebUI) loadPage(url string, w http.ResponseWriter) (*Page, error) {

	pageType, fields, err := a.parseURL(url)
	if err != nil {
		return &Page{Title: url, Body: []byte(err.Error())}, nil
	}

	if pageType == "home" {
		body := "Welcom to the flashback album  <br>"
		body = body + " We have " + strconv.Itoa(a.album.numPic) + " pictures in the album"
		return &Page{Title: url, Body: []byte(body)}, nil
	} else if pageType == "listYear" {
		yearInt, err := strconv.Atoi(fields["Year"])
		if err != nil {
			return &Page{Title: url, Body: []byte("Error processing Year")}, nil
		}
		body, _ := a.showPhotosYear(yearInt)
		return &Page{Title: url, Body: body}, nil
	} else if pageType == "listmonth" {
		//		yearint, err := strconv.atoi(fields["year"])
		//		monthint, err := strconv.atoi(fields["month"])
		return &Page{Title: url, Body: []byte("list year/month not implemented , pr are welcome :)")}, nil
	} else if pageType == "listPastYearSameWeekNum" {
		yearInt, err := strconv.Atoi(fields["Year"])
		if err != nil {
			return &Page{Title: url, Body: []byte("Error with the year number ")}, nil
		}
		monthInt, err := strconv.Atoi(fields["Month"])
		if err != nil {
			return &Page{Title: url, Body: []byte("Error with the Month number ")}, nil
		}

		dayInt, err := strconv.Atoi(fields["Day"])
		if err != nil {
			return &Page{Title: url, Body: []byte("Error with the day number ")}, nil
		}

		dayRequested := time.Date(yearInt, time.Month(monthInt), dayInt, 1, 50, 59, 0, time.UTC)
		body, _ := a.showWeekPhotos(dayRequested)
		if err != nil {
			return &Page{Title: url, Body: []byte("error retrieving photos")}, nil
		}

		return &Page{Title: url, Body: body}, nil
	}

	return &Page{Title: url, Body: []byte("Not expected for now")}, nil
}

// showPhotosYear : return a page with the list of photos for all years with this week number.
//
//	return []byte with the page content
func (a *AlbumWebUI) showWeekPhotos(dateSelected time.Time) ([]byte, error) {
	photos, _ := a.album.GetLstPhotosForWeek(dateSelected)

	body := "List picture of all year for the week : "
	for _, photo := range photos {
		body = body + string(a.showPhotoWithImg(photo)) + "<br>"
	}

	return []byte(body), nil

}

func (a *AlbumWebUI) showPhotoWithImg(pic *Photo) []byte {

	// TODO I think I need to use the option B from this site : https://zetcode.com/golang/http-serve-image/
	body := "<p>File: " + pic.photoPath + "</p> \n"
	body = body + "<a href=\"/" + pic.photoPath + "\">"
	body = body + "<img src=/" + pic.photoPath + " alt=\"" + pic.photoPath + "\" style=\"width:700px\">"
	body = body + "</a> <br> \n"
	body = body + pic.gps
	return []byte(body)
}

// showPhotosYear : return a page with the list of photos for the year in argument
//
//	return []byte with the page content
func (a *AlbumWebUI) showPhotosYear(year int) ([]byte, error) {

	photos, _ := a.album.GetLstPhotosForYear(year)

	body := "List year picture : <br>"
	for _, photo := range photos {
		var buf bytes.Buffer
		photo.PrintMainInfo(&buf)

		// <img src="pic_trulli.jpg" alt="Italian Trulli">
		body = body + "<img src=\"" + photo.photoPath + "\" alt=\"photo\" <br>"
		photo.PrintAllMetaData()

		body = body + buf.String() + "<br>"
	}

	return []byte(body), nil
}

// TODO review doc below

// parseURL , Parse & validate the URL provided and return what kind of page must be display
// return
//
//	      string			     indication which page will be display based on the URLRegex
//			 map[string]string	 return fields of the URLRegex
//			 error
func (a *AlbumWebUI) parseURL(url string) (string, map[string]string, error) {

	URLRegex := a.appPrefix + `/?(?P<Year>\d{4})?/?(?P<Month>\d{2})?/?(?P<Day>\d{2})?`
	r := regexp.MustCompile(URLRegex)

	regroup := FindAllGroups(r, url)

	if regroup == nil {
		return "", nil, fmt.Errorf("URL provided doesn't match RE : %s", URLRegex)
	}

	//	supportedPage: []string{"home", "listYear", "listMonth", "listDay"},
	// Check if the URL containt a day
	if regroup["Day"] != "" {
		day, err := strconv.Atoi(regroup["Day"])
		if day > 31 || err != nil {
			return "listPastYearSameWeekNum", regroup, fmt.Errorf("URL provided a Day bigger than 31, URL : %s", url)
		}
		return "listPastYearSameWeekNum", regroup, nil
	} else if regroup["Month"] != "" {
		month, err := strconv.Atoi(regroup["Month"])
		if month > 12 || err != nil {
			return "listMonth", regroup, fmt.Errorf("URL provided a Month bigger than 12, URL : %s", url)
		}
		return "listMonth", regroup, nil
	} else if regroup["Year"] != "" {
		return "listYear", regroup, nil
	} else if url == a.appPrefix || url == a.appPrefix+"/" {
		return "home", regroup, nil
	}

	return "", regroup, fmt.Errorf("URL not supported  : %s", url)
}

func (a *AlbumWebUI) viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path
	p, _ := a.loadPage(title, w)
	header := `
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Home page</title>
</head>
    `
	fmt.Fprintf(w, header+"<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

// LoadAlbums , provide a mecanismes to load pictures in the album
func (a *AlbumWebUI) LoadAlbums(logger *logrus.Entry) {
	logger.Debug("Load All pics:")
	numLoadedPic, err := a.album.LoadPhotosInAlbums()
	logger.Infof("We have loaded : %d , in directory : %s", numLoadedPic, album.directoryPaths)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

// FindAllGroups returns a map with each match group. The map key corresponds to the match group name.
// A nil return value indicates no matches.
func FindAllGroups(re *regexp.Regexp, s string) map[string]string {
	matches := re.FindStringSubmatch(s)
	subnames := re.SubexpNames()
	if matches == nil || len(matches) != len(subnames) {
		return nil
	}

	matchMap := map[string]string{}
	for i := 1; i < len(matches); i++ {
		matchMap[subnames[i]] = matches[i]
	}
	return matchMap
}
