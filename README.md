# screendream

Friendly neighborhood video streaming software with a focus on multi-user
playlist creation and fetching external videos. Support for playing local 
files is planned.

## Features (WIP):

- download and play (youtube) videos
  - adjust volume
  - skip through the video

- show a preview image for each video

- save and load playlists

- upload and play local files
	
-show nnd-like real time text messages on screen

## Installation

```sh
git clone https://github.com/sm0ken/screendream
cd screendream
```

To build and run the app you will need to install Go, mpv and youtube-dl

Currently it only depends on the Go standard library.

## Usage

```
go run src/github.com/sm0ken/screendream/main.go
```

Someone please tell me how to compile this to a binary	
