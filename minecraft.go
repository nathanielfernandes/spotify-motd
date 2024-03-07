package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/jpeg"
	"image/png"
	"net/http"

	"github.com/nathanielfernandes/motd/slp"
	"github.com/nfnt/resize"
)

func formatTime(milliseconds int) string {
	seconds := milliseconds / 1000
	minutes := seconds / 60
	seconds = seconds % 60
	return fmt.Sprintf("%d:%02d", minutes, seconds)
}

func ActivityToStatusResponse(activity Option[Spotify]) slp.StatusResponse {
	if activity := activity.Some(); activity != nil {
		artists := ""
		for i, artist := range activity.Artists {
			if i != 0 {
				artists += ", "
			}
			artists += artist
		}

		return slp.StatusResponse{
			Version: slp.Version{
				Name:     fmt.Sprintf("\u00a7f%s", formatTime(activity.Duration)),
				Protocol: 754,
			},
			Description: slp.Description{
				Text: fmt.Sprintf("\u00a7b%s\u00a7r\n%s", activity.Title, artists),
			},
			Favicon: getAlbumCoverAsFavicon(activity.AlbumCoverURL).UnwrapOr(""),
		}
	}

	return slp.StatusResponse{
		Version: slp.Version{
			Name:     "\u00a7foffline",
			Protocol: 754,
		},
		Description: slp.Description{
			Text: "Not listening to Spotify",
		},
	}
}

func getAlbumCoverAsFavicon(url string) Option[string] {
	resp, err := http.Get(url)
	if err != nil {
		return None[string]()
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return None[string]()
	}

	img, err := jpeg.Decode(resp.Body)
	if err != nil {
		return None[string]()
	}

	// resize and encode as png
	img = resize.Resize(64, 64, img, resize.Lanczos3)
	out := new(bytes.Buffer)
	err = png.Encode(out, img)
	if err != nil {
		return None[string]()
	}

	// encode as base64
	base64Str := base64.StdEncoding.EncodeToString(out.Bytes())
	return Some("data:image/png;base64," + base64Str)
}
