package main

import (
	"fmt"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type LocalClientFlow struct {
	Messages []string
	Index    int
}

func (d *LocalClientFlow) Init() error {
	d.Messages = append(d.Messages, "{\"msg\": \"hello\"}")
	d.Messages = append(d.Messages, "{\"msg\": \"hello2\"}")
	d.Messages = append(d.Messages, "{\"msg\": \"hello3\"}")
	d.Index = 0
	return nil
}

func (d *LocalClientFlow) NextMessage() string {
	if len(d.Messages) < d.Index {
		return ""
	}
	msg := d.Messages[d.Index]
	d.Index++
	return msg
}

func LocalClient(c *cli.Context) error {
	log.Infof("Starting local client")

	TargetHostname = TargetHostname

	log.Infof("Target host: %s", TargetHostname)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: TargetHostname, Path: "/ws"}

	wsc, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		if err == websocket.ErrBadHandshake {
			log.Errorf("%s: Bad Handshake Code: %s", TargetHostname, resp.Status)
		} else {
			log.Errorf("Failed to connect to %s: %s", TargetHostname, err.Error())
		}
		return fmt.Errorf("failed to dial: %s", err.Error())
	}
	defer wsc.Close()

	var is_running bool = true

	messages := new(LocalClientFlow)
	messages.Init()

	go func() {
		for is_running {
			_, message, err := wsc.ReadMessage()
			if err != nil {
				log.Infof("Failed to read message: %s", err.Error())
				return
			}
			log.Infof("Received: %s", message)
			nextMessage := messages.NextMessage()
			if nextMessage != "" {
				wsc.WriteMessage(1, []byte(nextMessage))
			}
		}
	}()

	err = wsc.WriteMessage(1, []byte(messages.NextMessage()))
	if err != nil {
		log.Errorf("Failed to write hello messages: %s", err.Error())
	}

	for {
		select {
		case <-interrupt:
			err := wsc.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Errorf("Failed to gracefully close connection: %s", err.Error())
				return err
			}
			is_running = false
		}
	}

	return nil
}
