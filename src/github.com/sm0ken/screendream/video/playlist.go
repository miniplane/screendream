package video

import (
	"fmt"
	"reflect"
)

type Playlist struct {
	Content []*V
}

func NewPlaylist() *Playlist {
	ret := new(Playlist)
	return ret
}

func (P *Playlist) AddVideo(video *V) {
	// add a new video at the end of the playlist

	P.Content = append(P.Content, video)
}

func (P *Playlist) RemoveVideo(video *V) {
	// delete a video from the playlist

	for i, v := range P.Content {
		if reflect.DeepEqual(v, video) {
			P.Content = append(P.Content[:i], P.Content[i+1:]...)
		}
	}
}

func (P *Playlist) SwapVideos(video1, video2 *V) {
	// swap two videos in the playlist

	for i, v := range P.Content {
		if reflect.DeepEqual(v, video1) {
			P.Content[i] = video2
		} else if reflect.DeepEqual(v, video2) {
			P.Content[i] = video1
		}
	}

}

func (p *Playlist) string() string {
	s := ""
	for _, element := range p.Content {
		s += s + element.string() + "\n"
	}

	return s
}
