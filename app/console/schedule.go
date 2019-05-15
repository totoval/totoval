package console

import "github.com/totoval/framework/cmd"

func Schedule(schedule *cmd.Schedule) {
	schedule.Command("say:hello-world hi,totoval").EverySecond()
}
