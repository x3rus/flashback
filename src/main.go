package main

import (
	"os"
	"strings"

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
	flashbackWebUI := NewAlbumWebUI(album, logrus)

	// Load picture in a go routine
	go flashbackWebUI.LoadAlbums(logrus)

	flashbackWebUI.RunWebSrv()

}
