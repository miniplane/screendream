package mvc

import(
	"github.com/sm0ken/screendream/video"
)

type Model interface{
	DownloadAndEnqueue(url string)
	Play(vid *video.V)
	SkipTo(time int)
	Pause()
	GetNowPlaying() 
	GetPosition() int
}
