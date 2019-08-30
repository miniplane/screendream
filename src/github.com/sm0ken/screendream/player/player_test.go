package player


import "testing"
import "time"
import v "github.com/sm0ken/screendream/video"





func TestNyancatPause(*testing.T){
	v := v.NewVideo("nc","/tmp/nc.mkv", 420)
	Play(v)
	time.Sleep(time.Second*5)
	Pause()
	time.Sleep(time.Second*5)
	Pause()
	time.Sleep(time.Second*5)
	Kill()
}
