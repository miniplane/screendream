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
var stdin io.WriteCloser
var observers []*Observer

//plays the video passed as vid
func Play(vid *v.V) {
	//kill other playing videos
	if playerstatus != dead && command != nil {
		Kill()
	}
	//start vid
	command = e.Command("mplayer", "-slave", vid.Location)
	playerstatus = playing
	
	//somebody(tm) should implement error handling
	stdin, _ = command.StdinPipe()
	command.Start()
}

// kills the currently playing video
func Kill() {
	command.Process.Kill()
	command = nil
}

func SkipTo(time int) {
	io.WriteString(stdin, "seek "+strconv.Itoa(time)+" 2\n")
}

func Pause(){
	io.WriteString(stdin, "pause\n")
}


func Time() int{ //time in seconds
	
}

//true if ended, false if killed
func Callback() bool{
	cb <- callback
	return cb
}


func waitForTermination(){
	if playerstatus!=dead{
		command.Wait()
		playerstatus=dead
	}
}


