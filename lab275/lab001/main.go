package main

import (
	"fmt"
	"github.com/jamesnetherton/m3u"
)

func main() {
	//playlist, err := m3u.Parse("https://raw.githubusercontent.com/jamesnetherton/m3u/master/testdata/playlist.m3u")
	playlist, err := m3u.Parse("https://raw.githubusercontent.com/Kimentanm/aptv/master/m3u/iptv.m3u")

	if err == nil {
		for _, track := range playlist.Tracks {
			fmt.Println("Track name:", track.Name)
			fmt.Println("Track length:", track.Length)
			fmt.Println("Track URI:", track.URI)
			fmt.Println("Track Tags:")
			for i := range track.Tags {
				fmt.Println(" -", track.Tags[i].Name, "=>", track.Tags[i].Value)

			}
			fmt.Println("----------")
		}
	} else {
		fmt.Println(err)
	}
}
