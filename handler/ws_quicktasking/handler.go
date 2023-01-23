package ws_quicktasking

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/avast/retry-go"
	"github.com/eagle/handler/auth"
	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/utils"
	"github.com/getsentry/sentry-go"
	"github.com/valyala/fastjson"
	"nhooyr.io/websocket"
)

func makeTLSConfig() *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: true,
	}
}

func makeTransport() *http.Transport {
	return &http.Transport{
		ForceAttemptHTTP2: true,
		TLSClientConfig:   makeTLSConfig(),
	}
}

func handleWebsocket(success chan bool) {
	defer sentry.Recover()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var err error
	var authed = false

	client := &http.Client{Timeout: 15 * time.Second}
	client.Transport = makeTransport()

	var c *websocket.Conn
	options := websocket.DialOptions{
		HTTPClient: client,
	}

	defer log.Fatalln("Tried to reconnect 10 times to websocket server, but failed. Closing bot...")

	_ = retry.Do(func() error {
		defer time.Sleep(1 * time.Second)
		// auth := loading.Data.Env.Env.AUTH_WEBSOCKET
		c, _, err = websocket.Dial(ctx, "wss://h90r03l2q4.execute-api.us-east-1.amazonaws.com/quicktask", &options)
		if err != nil {
			fmt.Println("Failed to connect to websocket server. Retrying...")
			return err
		} else {
			fmt.Println("Successfully connected to quicktask websocket.")
			fmt.Print("\033[H\033[2J")
			utils.Banner()
			auth.Welcome()
			logs.LogsMsgCyan("EagleBot is waiting for quicktasks...")
			// console.AddMQT()
		}

		for {
			_, message, err := c.Read(ctx)
			if err != nil {
				if errors.Is(err, websocket.CloseError{Code: websocket.StatusPolicyViolation, Reason: ""}) {
					log.Fatalln("Failed to authenticate to websocket server.")
				} else {
					log.Println("Error getting websocket message.")
					return err
				}
			}

			if !authed {
				if fastjson.GetBool(message, "success") {
					go func() { success <- true }()
					authed = true
				}
			}

			//check if the id_userKey is the same as the one in the request
			// id_userKey := auth.Auth.Integrations.Discord.ID

			if authed {
				if fastjson.Exists(message, "siteId") {
					go handleQuicktaskMessage(message)
				}
			}
		}
	}, retry.Attempts(10), retry.MaxDelay(15*time.Second), retry.RetryIf(func(err error) bool {
		return ctx.Err() == nil && !errors.Is(err, context.Canceled)
	}))
}
