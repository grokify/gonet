package main

import (
	"fmt"
	"log"

	"github.com/emersion/go-imap"
	"github.com/grokify/gonet/imaputil"
	"github.com/grokify/mogo/config"
)

func main() {
	err := config.LoadDotEnvSkipEmpty(".env")
	if err != nil {
		log.Fatal(err)
	}

	cm, err := imaputil.NewClientMoreEnv(imaputil.DefaultEnvPrefix)
	if err != nil {
		log.Fatal(err)
	}
	err = cm.ConnectAndLogin()
	if err != nil {
		log.Fatal(err)
	}
	defer cm.Logout()

	// List mailboxes
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- cm.Client.List("", "*", mailboxes)
	}()

	log.Println("Mailboxes:")
	for m := range mailboxes {
		log.Println("* " + m.Name)
	}

	fmt.Println("DONE")
}
