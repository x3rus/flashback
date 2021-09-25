package main

import (
	"fmt"

	exiftool "github.com/barasher/go-exiftool"
)

// Photo , structure to store pictures information
type Photo struct {
	photoPath string
	metadata  exiftool.FileMetadata
}

// NewPhoto , return Photo struct with nil for the metadata
func NewPhoto(filePath string) *Photo {
	return &Photo{
		photoPath: filePath,
	}
}

// LoadPhotoTags , Load metadata informations from the file on the harddrive
func (p *Photo) LoadPhotoTags() error {

	// Initialize exiftool
	et, err := exiftool.NewExiftool()
	if err != nil {
		return err
	}
	defer et.Close()

	// Extract Metadata, we can provide a list of photo
	fileInfos := et.ExtractMetadata(p.photoPath)
	if len(fileInfos) == 1 {
		//Extract the metadata information for the first file
		p.metadata = fileInfos[0]
	} else {
		return fmt.Errorf("Unable to load exif metadata for :%s", p.photoPath)
	}

	// return error from the exiftool for the pressed file
	if p.metadata.Err != nil {
		return p.metadata.Err
	}
	return nil

}

// GetPhotoTag , return the value for a tag or an empty string if the field do not exist
func (p *Photo) GetPhotoTag(tagName string) interface{} {

	if val, ok := p.metadata.Fields[tagName]; ok {
		return val.(string)
	}

	return ""
}

// PrintAllMetaData information this function will be use for the development I will remove it later
func (p *Photo) PrintAllMetaData() {

	if p.metadata.Err != nil {
		fmt.Printf("Error concerning %v: %v\n", p.metadata.File, p.metadata.Err)
	}
	for k, v := range p.metadata.Fields {
		fmt.Printf("[%v] %v\n", k, v)
	}

}
