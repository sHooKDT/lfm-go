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
	recent_track string = ""
	filepath string = "../media/music.txt"
	user string = "sHooKDT" // Example
	api_key string = "44995c1da3417933f9a380fde8342ff2" // My private key, dont use it if possibly (get new at last.fm/api)
	ApiURL string = fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=user.getRecentTracks&api_key=%s&user=%s&limit=5&format=xml", api_key, user)
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
	var trackdata = data.RecentTracks.Tracks[0]
	var total_string = fmt.Sprintf("%s - %s", trackdata.Name, trackdata.Artist)
	if total_string != recent_track {
		fmt.Printf("%s written. \n", total_string)
		recent_track = total_string
	}
	if trackdata.Nowplaying {
		return fmt.Sprintf(total_string)
	} else {return ""}
	
}

func main() {
	fmt.Println("Process started.")
	fmt.Printf("Writing in file: %s \n", filepath)
	for {
		ioutil.WriteFile(filepath, []byte(getTrack(ApiURL)), os.ModeTemporary)
		time.Sleep(2 * time.Second)
	}
}