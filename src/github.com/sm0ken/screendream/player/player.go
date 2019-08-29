package player

import (
	v "github.com/sm0ken/screendream/video"
	"io"
	e "os/exec"
	"strconv"
)

type status int

const (
	playing status = iota
	paused
	dead
)

var command *e.Cmd
var playerstatus status = dead

//plays the video passed as vid
func Play(vid *v.V) {
	//kill other playing videos
	if playerstatus != dead && command != nil {
		Kill()
	}
	//start vid
	command = e.Command("mplayer", "-slave", vid.Location)
	command.Start()
}

// kills the currently playing video
func Kill() {
	command.Process.Kill()
	command = nil
}

func SkipTo(time int) {
	if command != nil && playerstatus != dead {
		stdin, err := command.StdinPipe()
		defer stdin.Close()
		if err != nil {
			io.WriteString(stdin, "seek "+strconv.Itoa(time)+" 2\n")
		}
	}
}
