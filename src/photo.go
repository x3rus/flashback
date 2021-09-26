package main

import (
	"fmt"
	"reflect"
	"time"

	exiftool "github.com/barasher/go-exiftool"
)

// Photo , structure to store pictures information
type Photo struct {
	photoPath    string
	labels       []string
	gps          string // TODO review this type
	dateCreation time.Time
	fileType     string
	imageSize    string
	device       string
	metadata     exiftool.FileMetadata // Raw data extracted from the file
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

// SetPhotoStruct , get information from the raw data and set fields in the struct
func (p *Photo) SetPhotoStruct() error {

	var err error
	// TODO change to use the multiple error mechanism
	p.labels, err = p.extractPhotoLabel(p.metadata)
	p.gps, err = p.extractGPSPosition(p.metadata)
	p.dateCreation, err = p.extractDateCreation(p.metadata)
	p.fileType, err = p.extractFileType(p.metadata)
	p.imageSize, err = p.extractImageSize(p.metadata)
	p.device, err = p.extractDeviceUsed(p.metadata)
	return err
}

// extractPhotoLabel , extract from metadata the keywords label , tags set by the user
func (p *Photo) extractPhotoLabel(metadata exiftool.FileMetadata) ([]string, error) {

	// metadata can store the information in multiple format
	if val, ok := metadata.Fields["Keywords"]; ok {
		v := reflect.ValueOf(val)
		switch v.Kind() {
		case reflect.String:
			return []string{val.(string)}, nil
		case reflect.Slice:
			labels := make([]string, len(val.([]interface{})))
			for i, v := range val.([]interface{}) {
				labels[i] = fmt.Sprint(v)
			}
			return labels, nil
		default:
			return []string{""}, fmt.Errorf("Error type for field Keywords ,Found type:  %v", val)
		}
	}

	return []string{""}, nil
}

// extractGPSPosition, extract from metadata the keywords GPSPosition
func (p *Photo) extractGPSPosition(metadata exiftool.FileMetadata) (string, error) {

	// metadata can store the information in multiple format
	if val, ok := metadata.Fields["GPSPosition"]; ok {
		return val.(string), nil
	}

	return "", nil
}

// extractDateCreation, extract from metadata the keywords dateCreation
func (p *Photo) extractDateCreation(metadata exiftool.FileMetadata) (time.Time, error) {

	defaultTime := time.Date(1789, time.July, 14, 12, 0, 0, 0, time.UTC)
	// DateTimeOriginal look like this : 2016:07:24 16:59:32
	if val, ok := metadata.Fields["DateTimeOriginal"]; ok {
		// Writing down the way the standard time would look like formatted our way
		layout := "2006:01:02 15:04:05"
		t, err := time.Parse(layout, val.(string))
		if err != nil {
			return defaultTime, err
		}

		return t, nil
	}

	return defaultTime, nil
}

// extracFileType, extract from metadata the keywords FileType
func (p *Photo) extractFileType(metadata exiftool.FileMetadata) (string, error) {

	// metadata can store the information in multiple format
	if val, ok := metadata.Fields["FileType"]; ok {
		return val.(string), nil
	}

	return "", nil
}

// extracImageSize, extract from metadata the keywords ImageSize
func (p *Photo) extractImageSize(metadata exiftool.FileMetadata) (string, error) {

	// metadata can store the information in multiple format
	if val, ok := metadata.Fields["ImageSize"]; ok {
		return val.(string), nil
	}

	return "", nil
}

// extracDeviceUsed, extract from metadata the keywords DeviceType + Make
func (p *Photo) extractDeviceUsed(metadata exiftool.FileMetadata) (string, error) {

	var deviceUsed string
	if val, ok := metadata.Fields["DeviceType"]; ok {
		deviceUsed = val.(string)
	}

	if val, ok := metadata.Fields["Make"]; ok {
		deviceUsed = deviceUsed + ": " + val.(string)
	}
	return deviceUsed, nil
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

// PrintMainInfo , Print most important information about the picture
func (p *Photo) PrintMainInfo() {

	if p.metadata.Err != nil {
		fmt.Printf("Error concerning %v: %v\n", p.metadata.File, p.metadata.Err)
	}

	fmt.Printf(" File name : %s \n", p.photoPath)
	fmt.Printf(" 			date : %s \n", p.dateCreation)
	fmt.Printf(" 			labels : %s \n", p.labels)
	fmt.Printf(" 			location : %s \n", p.gps)
	fmt.Printf(" 			device used : %s \n", p.device)
	fmt.Printf(" 			image size: %s \n", p.imageSize)

}
