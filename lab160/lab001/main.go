package main

import (
	"encoding/json"
	"github.com/vansante/go-ffprobe"
	"log"
	"time"
)

func main() {
	path := "E:/Workspace/ffmpeg_labs/1-1导学.mp4"

	data, err := ffprobe.GetProbeData(path, 500*time.Millisecond)
	if err != nil {
		log.Panicf("Error getting data: %v", err)
	}

	buf, err := json.MarshalIndent(data, "", "  ")
	log.Println("ffprobe data:", string(buf))

	buf, err = json.MarshalIndent(data.GetFirstVideoStream(), "", "  ")
	log.Println("first video stream data:", string(buf))

	//log.Printf("\nDuration: %v\n", data.Format.Duration())
	//log.Printf("\nStartTime: %v\n", data.Format.StartTime())

	log.Println("时长:", data.Format.Duration())
	log.Println("时长(秒):", data.Format.DurationSeconds)
}
