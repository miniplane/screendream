package inmemorymodel


import(
	"github.com/sm0ken/screendream/downloader"
	"github.com/sm0ken/screendream/video"
)

type M struct{
	videos []*video.V
}


func NewM() *M{
	ret:=new(M)
	ret.videos = make([]*video.V,0,50)
}




func DownloadAndEnqueue(url string){
	// download video
	
}
