package video

import "strconv"

type V struct {
	Name     string
	Location string
	Duration int
}

func NewVideo(name string, loc string, dur int) *V {
	ret := new(V)
	ret.Name = name
	ret.Location = loc
	ret.Duration = dur
	return ret
}


func (v *V) string() string {
	s := v.Name + ", " + v.Location + ", " + strconv.Itoa(v.Duration)
	return (s)
}

