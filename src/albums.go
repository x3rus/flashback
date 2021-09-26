package main

// Albums Type  contain photos
type Albums struct {
	lstPhotos      []Photo
	directoryPaths []string
}

// NewAlbum , return Albums struct
func NewAlbum(directoryPaths []string) *Albums {
	return &Albums{
		directoryPaths: directoryPaths,
	}
}

// LoadPhotosInAlbums , iterate for each directory and load photos
func (a *Albums) LoadPhotosInAlbums() (int, error) {
	// Loop in each directory and load picture files
	return 0, nil
}
