package video


import (
	"reflect"
)



type Playlist struct {

	Content  []*V
	Length   int
}

func NewPlaylist() *Playlist {
	ret := new(Playlist)
	ret.Length = 0
	return ret
}

func (P *Playlist) AddVideo(video *V)  {
	P.Content = append(P.Content, video)
	P.Length++
}

func (P *Playlist) RemoveVideo(video *V) {

	for i, v := range P.Content {
		if reflect.DeepEqual(v, video) {
			P.Content = append(P.Content[:i], P.Content[i+1:]...)
		}
	}
}


func (p *Playlist) string() string {
	s := ""
	for _, element := range p.Content {
		s = s + element.string() + "\n"
	}

	return (s)
}


