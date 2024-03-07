package main

import (
	"fmt"
	"net"

	"github.com/nathanielfernandes/motd/slp"
)

// singletons, my beloved
var activity Option[Spotify] = None[Spotify]()
var activityResponse slp.StatusResponse = ActivityToStatusResponse(None[Spotify]())

func main() {
	// Listen for Spotify activity changes
	go func() {
		activites := ListenForSpotify()
		for {
			activity = <-activites
			activityResponse = ActivityToStatusResponse(activity)
		}
	}()

	// Listen for Minecraft connections
	ln, err := net.Listen("tcp", ":25565")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fmt.Println("Server is running on port 25565")

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go slp.HandleConnection(conn, status, disconnect)
	}

}

func status(_ net.Conn) slp.StatusResponse {
	return activityResponse
}

func disconnect(_ net.Conn, __ slp.LoginStart) slp.Disconnect {
	if ac := activity.Some(); ac != nil {
		return slp.DisconnectWithStringMsg(ac.TrackUrl)
	}
	return slp.DisconnectWithStringMsg("This is not a real Minecraft server.")
}
