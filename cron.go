package inhabot

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
)

var scheduler *gocron.Scheduler

func Cron(discord *discordgo.Session, Titles *formertitlelist) *gocron.Scheduler {
	Scheduler := gocron.NewScheduler(time.Local)
	Scheduler.SetMaxConcurrentJobs(3, gocron.WaitMode)
	Scheduler.Cron("0 0/1 * * *").Do(CheckUpdate, discord, CSE)
	return Scheduler
}

/*
func CheckTime() {
	loc, _ := time.LoadLocation("Asia/Seoul")
	log.Println("Scheduler works at ", time.Now().In(loc))
}
*/
