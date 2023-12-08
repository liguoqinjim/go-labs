package main

import (
	"fmt"
	"github.com/AlekSi/pointer"
	"github.com/quangngotan95/go-m3u8/m3u8"
)

func main() {
	playlist := m3u8.NewPlaylist()

	item := &m3u8.PlaylistItem{
		Width:      pointer.ToInt(1920),
		Height:     pointer.ToInt(1080),
		Profile:    pointer.ToString("high"),
		Level:      pointer.ToString("4.1"),
		AudioCodec: pointer.ToString("aac-lc"),
		Bandwidth:  540,
		URI:        "test.url",
	}
	playlist.AppendItem(item)

	item2 := &m3u8.MediaItem{
		Type:          "AUDIO",
		GroupID:       "audio-lo",
		Name:          "Francais",
		Language:      pointer.ToString("fre"),
		AssocLanguage: pointer.ToString("spoken"),
		AutoSelect:    pointer.ToBool(true),
		Default:       pointer.ToBool(false),
		Forced:        pointer.ToBool(true),
		URI:           pointer.ToString("frelo/prog_index.m3u8"),
	}
	playlist.AppendItem(item2)

	fmt.Println(playlist.String())
}
