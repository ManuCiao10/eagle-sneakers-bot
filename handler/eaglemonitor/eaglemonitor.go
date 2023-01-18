package eaglemonitor

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/quicktask"
	"github.com/eagle/handler/task_manager"
	"github.com/eagle/handler/utils"
)

var (
	run       = false
	allPidMqt = []string{}
)

func monitorInitialize() {
	Token := loading.Data.Env.Env.TBAtoken

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)
	// we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	logs.LogsMsgCyan("EagleBot is now monitoring...")

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself, it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	for _, dataDiscord := range m.Embeds {
		pid := getPid(dataDiscord)
		store := getSite(dataDiscord)
		size := "random"

		for _, taskUUID := range loading.Data.Quicktask.Quicktask[store] {
			taskObject, err := quicktask.GetQuicktask(taskUUID)

			if err != nil {
				fmt.Println("Failed to get task: ", err.Error())
				continue
			}

			delay := loading.Data.Settings.Settings.Delay.Retry
			delayInt, err := strconv.Atoi(delay)
			if err != nil {
				fmt.Println("Failed to convert delay to int: ", err.Error())
				continue
			}
			quantity, err := strconv.Atoi(taskObject.Tasks_Quantity)
			if err != nil {
				fmt.Println("Failed to convert quantity to int: ", err.Error())
				continue
			}
			pidMqt := strings.Split(strings.ToLower(taskObject.Other), ";")

			allPidMqt = append(allPidMqt, pidMqt...)
			taskObject.Pid = pid

			if utils.ContainsPID(allPidMqt, pid) {
				if !run {
					run = true
					logs.LogsMsgCyan("restock detected!")
					logs.MonitorWebhook(&logs.MonitorDetected{
						Pid:          pid,
						Size:         size,
						TaskQuantity: quantity,
						Proxy:        taskObject.Proxylist,
						TaskFile:     taskObject.Accounts,
						Delay:        delayInt,
						Store:        store,
					}, loading.Data.Settings.Settings.DiscordWebhook)
				}

				if !taskObject.Active {
					go task_manager.RunQuickTask(taskObject)
				} else if taskObject.Done {
					task_manager.StopQuickTask(taskObject)
				}

			}
		}
		time.Sleep(1 * time.Millisecond)
	}

}
