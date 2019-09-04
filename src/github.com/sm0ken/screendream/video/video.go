package video

import (
	"log"
	"net/url"
	"strconv"
)

type V struct {
	Name     string
	Location string
	Duration int
	URL      url.URL
}

func NewVideo(name string, loc string, dur int, rawurl string) *V {
	ret := new(V)
	ret.Name = name
	ret.Location = loc
	ret.Duration = dur
	uri, err = url.Parse(rawurl)

	if err != nil {
		log.Fatal(err)
	} else {
		ret.URL = uri
	}

	return ret
}

func (v *V) string() string {
	s := v.Name + ", " + v.Location + ", " + strconv.Itoa(v.Duration)
	return (s)
}
