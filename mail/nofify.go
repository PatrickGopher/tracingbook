package mail

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
	"tracingbook/models"
)

//type Notify interface {
//	Notify() bool
//}

type NotifyByMail struct {
	To      string
	From    string
	Content string
	Title   string
	Pass    string
}

func (nm *NotifyByMail) NotifyUpdates(updateItems []models.UpdateItem) {
	if len(updateItems) == 0 {
		return
	}
	updates := make([]string, len(updateItems))
	title := updateItems[0].BookName + " 最新更新:" + updateItems[0].LatestName
	for i, item := range updateItems {
		update := fmt.Sprintf("最新章节:%s, 点击: %s \n", item.LatestName, item.BookUrl)
		updates[i] = update
	}
	NML.Content = strings.Join(updates, "\n")
	if len(title) > 200 {
		title = title[0:200]
	}
	NML.Title = title
	NML.notify()
}

func (nm *NotifyByMail) notify() bool {
	msg := "From: " + nm.From + "\n" +
		"To: " + nm.To + "\n" +
		"Subject:" + nm.Title + "\n\n" +
		nm.Content
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", nm.From, nm.Pass, "smtp.gmail.com"),
		nm.From, []string{nm.To}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return false
	}

	return true
}

var NML = NotifyByMail{
	From:    "mixius.life@gmail.com",
	To:      "zhe.chen.sg@gmail.com",
	Content: "test test",
	Title:   "test from go",
	Pass:    "",
}
