package main

import "net/http"

type FlashMessage struct {
	Type string
	Message string
}

func getFlashMessages(w http.ResponseWriter, r *http.Request) []FlashMessage{
	session, err := store.Get(r, "front-end")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return []FlashMessage{}
	}

	if flashes := session.Flashes(); len(flashes) > 0 {
		//todo parse messages
	}
	//return []FlashMessage{{Message:"test", Type: "alert"}}
	return []FlashMessage{}
}