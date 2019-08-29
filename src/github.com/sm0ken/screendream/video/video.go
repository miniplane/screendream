package video

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
