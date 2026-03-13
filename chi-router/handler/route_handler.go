package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandlerHomeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is a home route")
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	// get id params
	id := chi.URLParam(r, "id")
	fmt.Println("The value of id is:", id)
}

func HandleUserChat(w http.ResponseWriter, r *http.Request) {
	// get id and chat id params
	id := chi.URLParam(r, "id")
	chatId := chi.URLParam(r, "chatid")

	fmt.Printf("The user id is %v and chat id is %v.\n", id, chatId)
}
