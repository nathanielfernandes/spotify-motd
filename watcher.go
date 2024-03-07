package main

import (
	"encoding/json"

	"github.com/donovanhide/eventsource"
)

type Spotify struct {
	Album         string   `json:"album"`
	AlbumCoverURL string   `json:"album_cover_url"`
	AlbumShareUrl string   `json:"album_share_url"`
	Artist        string   `json:"artist"`
	Artists       []string `json:"artists"`
	Title         string   `json:"title"`
	TrackId       string   `json:"track_id"`
	TrackUrl      string   `json:"track_url"`
	Start         int      `json:"start"`
	End           int      `json:"end"`
	Duration      int      `json:"duration"`
}

type Activity struct {
	Type         string  `json:"type"`
	ActivityData Spotify `json:"activity"`
}

// server sent events endpoint
const live_activies = "https://watcher.ncp.nathanferns.xyz/live-activity/485138947115057162"

func ListenForSpotify() chan Option[Spotify] {
	activityChan := make(chan Option[Spotify])

	// connect to the server sent events endpoint
	stream, err := eventsource.Subscribe(live_activies, "")
	if err != nil {
		panic(err)
	}

	// listen for events
	go func() {
		for {
			select {
			case event := <-stream.Events:
				activityChan <- parseEventData(event.Data())

			case err := <-stream.Errors:
				panic(err)
			}
		}
	}()

	return activityChan
}

func parseEventData(data string) Option[Spotify] {
	var activities []Activity
	err := json.Unmarshal([]byte(data), &activities)
	if err != nil || len(activities) == 0 {
		return None[Spotify]()
	}

	spotify := None[Spotify]()
	for _, activity := range activities {
		if activity.Type == "Spotify" {
			spotify = Some(activity.ActivityData)
			break
		}
	}

	return spotify
}
