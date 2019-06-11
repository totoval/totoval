package jobs

import "github.com/totoval/framework/job"

func Initialize() {
	// initialize topic and channel
	job.RegisterQueue()
}
