package inmemorymodel


import(
	"github.com/sm0ken/screendream/downloader"
	"github.com/sm0ken/screendream/video"
	"sync"
)

type M struct{
	videos []*video.V
	listMux *sync.Mutex
}


func NewM() *M{
	ret:=new(M)
	ret.videos = make([]*video.V,0,50)
	listMux = new(sync.Mutex)
}



//calls worker method in goroutine
func DownloadAndEnqueue(url string){
	go downloadAndEnqueue(string)
}



func downloadAndEnqueue(url string){
	// download video
	v := downloader.Download_video(url)
	listMux.Lock()
	append(videos, v)
	listMux.Unlock()
}
