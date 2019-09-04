package downloader

import (
	"fmt"
	"log"
	"os/exec"
	"os"
	"strings"
	"strconv"
	"github.com/sm0ken/screendream/video"
)

func Download_video(url string) *video.V {
	// download a video using youtube-dl
	// save video to temp
	// return video struct containing metadata of the video

	cmd := exec.Command("youtube-dl", url)

	if ! exists("/tmp/screendream") {
		create_folder("/tmp/screendream")
	}

	cmd.Dir = "/tmp/screendream"

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Video loaded.")

	title 	 := get_metadata(url, "--get-title")

	duration := get_metadata(url, "--get-duration")
	seconds  := parse_duration(duration)

	location := get_metadata(url, "--get-filename")
	location = "/tmp/screendream/"+location


	vid := video.NewVideo(title, location, url, seconds)
	
	return vid
}


func parse_duration(duration string) int {
	// parse a string containing the video duration in the format mm:ss
	// convert minutes to seconds
	// doesn't work for hour long videos yet

	err1, err2, err3 := error(nil), error(nil), error(nil)
	seconds, minutes, hours := 0, 0, 0

	s := strings.Split(duration, ":")
	seconds, err1 = strconv.Atoi(s[len(s)-1])

	if len(s) >= 2 {

		minutes, err2 = strconv.Atoi(s[len(s)-2])

		if len(s) == 3 {
			hours, err3 = strconv.Atoi(s[0])
		}
	}

	if err1 != nil || err2 != nil ||  err3 != nil {
        fmt.Println(err1, err2, err3)
        // todo log
    }

	return seconds + minutes*60 + hours*60*60

}


func create_folder(name string) {
	// create a new folder

	os.Mkdir(name, 0777)
}


func exists(path string) (bool) {
	// check if folder path exists

    _, err := os.Stat(path)
    if err == nil { return true }
    if os.IsNotExist(err) { return false }
    return true
}


func get_metadata(url string, option string) string {
	// use youtube-dl to extract metadata needed for video struct

	out, err := exec.Command("youtube-dl", option, url).Output()

	if err != nil {
		log.Fatal(err)
	}

	output := string(out[:])
	return strings.Trim(output, "\n")

}

/*
func main() {
	download_video("https://www.youtube.com/watch?v=1IemshjOnLE")
}
*/
