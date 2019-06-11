package listeners

import "github.com/totoval/framework/hub"

func Initialize() {
	// initialize topic and channel
	hub.RegisterQueue()
}
