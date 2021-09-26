package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

// Page struct for the web server
type Page struct {
	Title string
	Body  []byte
}

// AlbumWebUI web interface for the UI
type AlbumWebUI struct {
	album  *Albums
	logger *logrus.Entry
}

// NewAlbumWebUI , return Albums struct
func NewAlbumWebUI(album *Albums, logger *logrus.Entry) *AlbumWebUI {
	return &AlbumWebUI{
		album:  album,
		logger: logger,
	}
}

// RunWebSrv , provide the web interface
func (a *AlbumWebUI) RunWebSrv() {
	http.HandleFunc("/", a.viewHandler)
	a.logger.Fatal(http.ListenAndServe(":8080", nil))
}

func (a *AlbumWebUI) loadPage(url string) (*Page, error) {

	if url == "/" {
		body := "Welcom to the flashback album  <br>"
		body = body + " We have " + strconv.Itoa(a.album.numPic) + " pictures in the album"
		return &Page{Title: url, Body: []byte(body)}, nil
	}

	return &Page{Title: url, Body: []byte("Not expected for now")}, nil
}

func (a *AlbumWebUI) viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path
	p, _ := a.loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
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
