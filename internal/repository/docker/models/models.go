package models

import "fmt"

// Image contains all necessary fields for downloading and running image
type Image struct {
	Title   string
	Tag     string
	Host    string
	FileExt string
	Cmd     func(filename string) []string
}

// ImageWithHost return string in "host/title:tag" format
func (i *Image) ImageWithHost() string {
	return fmt.Sprintf("%s/%s:%s", i.Host, i.Title, i.Tag)
}

// Image returns string in "title:tag" format
func (i *Image) Image() string {
	return fmt.Sprintf("%s:%s", i.Title, i.Tag)
}

// Info contains information required for running container
type Info struct {
	Language string
	Code     string
}
