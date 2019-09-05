package inmemorymodel


import(
	"github.com/sm0ken/screendream/downloader"
	v "github.com/sm0ken/screendream/video"
	"reflect"
	"sync"
	p "github.com/sm0ken/screendream/player"
)

var {
	videos []*video.V
	nowplaying *video.V
	listMux *sync.Mutex
}


func init(){
	videos = make([]*video.V,0,50)
	listMux = new(sync.Mutex)
}


func playlistPosition() int{
	listMux.Lock()
	for i:=0, i<len(videos),i++{
		if reflect.DeepEqual(nowplaying, videos[i]){
			return i
		}
	}
	return -1 //playing video has been removed --we should prevent this from happening if at all possible
}


//calls worker method in goroutine
func DownloadAndEnqueue(url string){
	go downloadAndEnqueue(string)
}



func downloadAndEnqueue(url string){
	// download video
	v := downloader.Download_video(url)
	if positionOf(v)!=-1{
		return
	}
	listMux.Lock()
	append(videos, v)
	listMux.Unlock()
}

func Play(vid *v.V){
	p.Play(vid)
}

func PlayerNotify(reason bool){
	if reason==true{ //ended normally
		//play next video
		pos:=playlistPosition()
		if pos ==-1{
			pos=len(videos)
		}else{
			pos++
		}
		p.Kill()
		p.Play(videos[pos])
		
	}else{ //killed
		
	}
}




