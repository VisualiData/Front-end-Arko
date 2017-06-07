package main

import (
	"net/http"
)

type FlashMessage struct {
	Type string
	Message string
}

func getFlashMessages(w http.ResponseWriter, r *http.Request) []FlashMessage{
	session, err := store.Get(r, "front-end")
	var flashMessages []FlashMessage
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return []FlashMessage{}
	}

	if flashes := session.Flashes(); len(flashes) > 0 {
		for i:=0; i < len(flashes); i++ {
			message := flashes[i].(*FlashMessage)
			flashMessages = append(flashMessages, FlashMessage{Message:message.Message, Type: message.Type})
		}
	}
	session.Save(r, w)
	return flashMessages
}

func addFlashMessage(w http.ResponseWriter, r *http.Request, f FlashMessage) {
	session, err := store.Get(r, "front-end")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.AddFlash(f)
	session.Save(r, w)
}