package video

import (
	"strconv"
)

type V struct {
	Name     string
	Location string
	Url 	 string
	Duration int
}

func NewVideo(name string, loc string, url string, dur int) *V {
	ret := new(V)
	ret.Name = name
	ret.Location = loc
	ret.Url = url
	ret.Duration = dur


	return ret
}

func (v *V) string() string {
	s := v.Name + ", " + v.Location + ", " + v.Url + ", " + strconv.Itoa(v.Duration)
	return (s)
}
