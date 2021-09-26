package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

// global variable to make it available in the web service
var album *Albums

// getEnv retreieve environnement variable or a default value.
func getenv(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}

// Page struct for the web server
type Page struct {
	Title string
	Body  []byte
}

func main() {

	// Logger
	logrus := log.NewEntry(log.New())

	// retrieve variable environnement and define default value if not present
	envAlbumDirs := getenv("ALBUMDIRS", "./data/pic-sample/dir1,./data/pic-sample/dir1")
	albumDirs := strings.Split(envAlbumDirs, ",")

	if len(albumDirs) < 1 {
		logrus.Errorf("ERROR: The album directory is empty: %s - please set Env variable ALBUMSDIRS", envAlbumDirs)
		os.Exit(1)
	}

	album = NewAlbum(albumDirs)
	go loadAlbums(logrus)

	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loadPage(url string) (*Page, error) {

	if url == "/" {
		body := "Welcom to the flashback album  <br>"
		body = body + " We have " + strconv.Itoa(album.numPic) + " pictures in the album"
		return &Page{Title: url, Body: []byte(body)}, nil
	}

	return &Page{Title: url, Body: []byte("Not expected for now")}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func loadAlbums(logger *logrus.Entry) {
	logger.Debug("Load All pics:")
	numLoadedPic, err := album.LoadPhotosInAlbums()
	logger.Infof("We have loaded : %d , in directory : %s", numLoadedPic, album.directoryPaths)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
