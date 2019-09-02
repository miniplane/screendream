package video

import (
	"testing"
	"fmt"
	"strconv"
)


func TestPlaylist(t *testing.T) {

	playlist := NewPlaylist()

	vid1 := NewVideo("asdf", "asdf", 46)
	playlist.AddVideo(vid1)

	vid2 := NewVideo("asdfa", "asdf", 46)
	playlist.AddVideo(vid2)

	vid3 := NewVideo("asdfb", "asdf", 46)
	playlist.AddVideo(vid3)

	fmt.Println("Playlist Length: "+strconv.Itoa(playlist.Length))
	fmt.Println("Current Playlist:\n"+playlist.string())

	playlist.RemoveVideo(vid2)

	fmt.Println("Playlist Length: "+strconv.Itoa(playlist.Length))
	fmt.Println("Current Playlist:\n"+playlist.string())

	playlist.RemoveVideo(vid1)
	playlist.RemoveVideo(vid3)

	fmt.Println("Playlist Length: "+strconv.Itoa(playlist.Length))
	fmt.Println("Current Playlist:\n"+playlist.string())

}


func TestVideoPrint(t *testing.T) {

	vid := NewVideo("asdf", "asdf", 46)

	fmt.Println("test:\n"+vid.string())

}

