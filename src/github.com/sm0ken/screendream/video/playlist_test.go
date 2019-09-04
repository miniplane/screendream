package video

import (
	"fmt"
	"testing"
)

func TestPlaylist(t *testing.T) {

	playlist := NewPlaylist()

	// adding videos
	vid1 := NewVideo("asdf", "asdf", 46)
	playlist.AddVideo(vid1)

	vid2 := NewVideo("asdfa", "asdf", 46)
	playlist.AddVideo(vid2)

	vid3 := NewVideo("asdfb", "asdf", 46)
	playlist.AddVideo(vid3)

	fmt.Println("Current Playlist:\n" + playlist.String())

	// swapping videos
	playlist.SwapVideos(vid1, vid2)

	fmt.Println("Current Playlist:\n" + playlist.String())

	// deleting videos
	playlist.RemoveVideo(vid2)

	fmt.Println("Current Playlist:\n" + playlist.String())

	playlist.RemoveVideo(vid1)
	playlist.RemoveVideo(vid3)

	fmt.Println("Current Playlist:\n" + playlist.String())

}

func TestVideoPrint(t *testing.T) {

	vid := NewVideo("asdf", "asdf", 46)

	fmt.Println("test:\n" + vid.string())

}
