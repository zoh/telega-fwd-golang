package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/zelenin/go-tdlib/client"
	lib "github.com/zoh/telega-fwd-golang/lib"
)

func main() {
	tdlibClient, err := lib.CreateClient()
	if err != nil {
		log.Fatalf("NewClient error: %s", err)
	}

	//bot := createBot()

	me, err := tdlibClient.GetMe()
	lib.CheckErr(err)

	log.Printf("Hello @%s!", me.Username)

	chats, err := tdlibClient.GetChats(
		&client.GetChatsRequest{
			OffsetOrder:  9223372036854775807,
			OffsetChatId: 0,
			Limit:        100,
		})
	if err != nil {
		log.Fatalf("getChats error: %s", err)
	}

	log.Println("Start find chats")
	privateOutcome := strings.Replace(os.Getenv("CHANNEL_PRIVATE_OUTCOME"), "'", "", -1)
	if privateOutcome == "" {
		panic("CHANNEL_PRIVATE_OUTCOME not defined!")
	}

	privateTarget := strings.Replace(os.Getenv("CHANNEL_TARGET"), "'", "", -1)
	if privateTarget == "" {
		panic("CHANNEL_TARGET not defined!")
	}

	log.Println(chats, "==>")
	// todo: and another
	var privateOutcomeID, privateTargetID int64
	for _, v := range chats.ChatIds {
		ch, err := tdlibClient.GetChat(&client.GetChatRequest{ChatId: v})
		lib.CheckErr(err)
		//
		if ch.Title == privateOutcome {
			privateOutcomeID = v
		}
		if ch.Title == privateTarget {
			privateTargetID = v
		}
	}

	if privateOutcomeID == 0 {
		panic("Не нашли CHANNEL_PRIVATE_OUTCOME в списке достуных каналов!")
	}
	if privateTargetID == 0 {
		panic("Не нашли CHANNEL_TARGET в списке достуных каналов!")
	}

	time.Sleep(time.Second)

	r := &client.SendMessageRequest{
		ChatId: privateOutcomeID,
		InputMessageContent: &client.InputMessageText{
			Text: &client.FormattedText{
				Text:     "test XUY!!! asd asd asd",
				Entities: nil,
			},
		},
	}
	mess, err := tdlibClient.SendMessage(r)
	log.Println(privateOutcome, privateOutcomeID, mess, err)

	log.Println("send!")
	// msg := tgbotapi.NewPhotoUpload(privateOutcomeID, "/vagrant/img.png")
	// msg.Caption = "Test"

	// _, err = bot.Send(msg)
	// checkErr(err)

	listener := tdlibClient.GetListener()
	defer listener.Close()

	for update := range listener.Updates {
		if update.GetClass() == client.ClassUpdate {
			if update.GetType() == "updateNewMessage" {
				data := update.(*client.UpdateNewMessage)

				log.Println("Content type:", data.Message.Content.MessageContentType())

				if data.Message.ChatId == privateTargetID {

					// if "messageText" == data.Message.Content.MessageContentType() {
					// 	msg := data.Message.Content.(*client.MessageText)

					// 	log.Println(msg.Text)
					// }

					fwr := &client.ForwardMessagesRequest{
						FromChatId: data.Message.ChatId,
						ChatId:     privateOutcomeID,
						MessageIds: []int64{data.Message.Id},
					}
					m, err := tdlibClient.ForwardMessages(fwr)
					if err != nil {
						log.Println("error", err)
					} else {
						log.Println("new messages", m.Messages)
					}
				}

			}
		}
	}
}
