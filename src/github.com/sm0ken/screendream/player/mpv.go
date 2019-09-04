package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strconv"
	"time"
)

type mpvTime struct {
	Data  float64
	Error string
}

type mpvPlay struct {
	Data  bool
	Error string
}

// plays a file at a given location
// not sure about how we want to pass file specifiers here
// just fooken yeet the link to the video in here alright
func Play(arg string) {
	app := "mpv"

	arg0 := "--input-ipc-server=/tmp/mpvsocket"
	arg1 := "--fs"

	cmd := exec.Command(app, arg0, arg1, arg)
	log.Printf("Playing %s in mpv, IPC enabled via /tmp/mpvsocket", arg1)

	out, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
		fmt.Print(out)
		return
	}
	log.Println("Successfully finished playing %s", arg1)
}

func connect() net.Conn {
	c, err := net.Dial("unix", "/tmp/mpvsocket")
	if err != nil {
		log.Fatal(err)
		panic(err) // TODO: better error handling
	}

	return c
}

// does not use makeCmd because the value may be a literal such as a boolean
func setProperty(p, v string) string {
	return fmt.Sprintf(`{"command": ["set_property", "%s", %s]}`, p, v) + "\n"
}

func getProperty(p string) string {
	return makeCmd("get_property", p)
}

func execCmd(cmd string) {
	c := connect()
	defer c.Close()

	i, err := c.Write([]byte(cmd))

	if err != nil {
		log.Fatal(err)
		log.Fatal(fmt.Sprintf("Command length was %d, wrote %d bytes.", len(cmd), i))
	}
}

func execRead(cmd string) (int, *[]byte) {
	c := connect()
	defer c.Close()

	i, err := c.Write([]byte(cmd))

	if err != nil {
		log.Fatal(err)
		log.Fatal(fmt.Sprintf("Command length was %d, wrote %d bytes.", len(cmd), i))
	}

	buf := make([]byte, 512)

	i, err = c.Read(buf)

	if err != nil {
		log.Fatal(err)
		log.Fatal(fmt.Sprintf("Error while reading communication from mpv. %d bytes read."), i)
	}

	return i, &buf
}

// exits mpv
func Quit() {
	execCmd(makeCmd("quit"))
}

// Currently broken, TODO either use static variable or get current state
func Toggle() {
	var cmd string

	i, buf := execRead(getProperty("pause"))

	var m mpvPlay

	err := json.Unmarshal((*buf)[:i], &m)

	if err != nil {
		log.Println("Error occured while unmarshaling:")
		log.Println(err)
	}

	if m.Data {
		cmd = setProperty("pause", "false")
	} else {
		cmd = setProperty("pause", "true")
	}

	execCmd(cmd)
}

func Pause() {
	execCmd(setProperty("pause", "false"))
}

func Resume() {
	execCmd(setProperty("pause", "false"))
}

func Jump(pos int) {
	execCmd(makeCmd("seek", strconv.Itoa(pos), "absolute"))
}

// returns the position in the video in seconds
func Position() int {
	i, buf := execRead(getProperty("playback-time"))

	var m mpvTime

	err := json.Unmarshal((*buf)[:i], &m)

	if err != nil {
		log.Println("Error occured:")
		log.Println(err)
	}

	return int(m.Data)
}

func makeCmd(args ...string) string {
	var res string

	for _, arg := range args {
		res += `"` + arg + `", `
	}

	return `{"command": [` + res[:(len(res)-2)] + `]}` + "\n"
}

func cycle(prop string) {
	c := connect()
	defer c.Close()

	cmd := makeCmd("cycle", prop)

	_, _ = c.Write([]byte(cmd))
}

func CycleAudio() {
	cycle("aid")
}

func CycleSubs() {
	cycle("sid")
}

func main() {
	go Play("/Users/zero/magus.mkv")

	time.Sleep(3000 * time.Millisecond)

	CycleAudio()
	CycleSubs()

	time.Sleep(1000 * time.Millisecond)

	Toggle()

	time.Sleep(5000 * time.Millisecond)
	Position()

	Jump(360)

	time.Sleep(2000 * time.Millisecond)
	Position()

	Quit()
}

// TODO: command for subtitles/audio tracks, anything else?
