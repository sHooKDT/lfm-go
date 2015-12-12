package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"os"
	"time"
)

var (
	ApiURL string = "http://ws.audioscrobbler.com/2.0/?method=user.getRecentTracks&api_key=44995c1da3417933f9a380fde8342ff2&user=sHooKDT&limit=5&format=xml"
)

type XMLTrack struct {
	XMLName xml.Name `xml:"track"`
	Nowplaying bool `xml:"nowplaying,attr"`
	Artist string `xml:"artist"`
	Name string `xml:"name"`
	// imageurl string `xml:"image"`
}

type XMLRecentTracks struct {
	XMLName xml.Name `xml:"recenttracks"`
	Tracks []XMLTrack `xml:"track"`
}


type XMLlfm struct {
	XMLName xml.Name `xml:"lfm"`
	RecentTracks XMLRecentTracks `xml:"recenttracks"`
}

func getTrack(URL string) string {
	response, _ := http.Get(ApiURL)
	XMLData, _ := ioutil.ReadAll(response.Body)
	var data XMLlfm
	err := xml.Unmarshal(XMLData, &data)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	var trackString = data.RecentTracks.Tracks[0].Name + " - " +  data.RecentTracks.Tracks[0].Artist
	return trackString
}

func main() {
	for {
		ioutil.WriteFile("track.txt", []byte(getTrack(ApiURL)), os.ModeTemporary)
		time.Sleep(2 * time.Second)
	}
}