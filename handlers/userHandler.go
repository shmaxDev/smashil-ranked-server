package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"smashil-ranked/dtos"
	internalErrors "smashil-ranked/errors"
	queueloop "smashil-ranked/queueLoop"
	"smashil-ranked/services"
	"time"

	"fmt"

	"github.com/go-playground/validator/v10"

	"github.com/gorilla/websocket"
)

var validate *validator.Validate

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(u *services.UserService) *UserHandler {
	return &UserHandler{u}
}

func (h *UserHandler) HandlePostPlayer(w http.ResponseWriter, r *http.Request) {
	var user dtos.UserDto

	err := DecodeAndValidate(&user, &w, r)
	if err != nil {
		return
	}

	_, err = h.UserService.AddUser(*user.Id)

	if err != nil {
		var httpErr *internalErrors.HTTPError

		if ok := errors.As(err, &httpErr); ok {
			http.Error(w, httpErr.Error(), httpErr.StatusCode)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) HandleAddToQueue(w http.ResponseWriter, r *http.Request) {
	var nig dtos.Queue

	err := json.NewDecoder(r.Body).Decode(&nig)

	if err != nil {
		if errors.Is(err, io.EOF) {
			http.Error(w, "Request body is empty", http.StatusBadRequest)
			return
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var player = queueloop.Player{UserId: *nig.UserId, Elo: *nig.Elo, TimeJoined: time.Now()}

	queueloop.Add(player)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var BotConn *websocket.Conn

func HandleBotSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}
	fmt.Println("Bot connected via WebSocket")
	BotConn = conn

	for {
		var msg map[string]any
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Bot disconnected:", err)
			BotConn.Close()
			BotConn = nil
			break
		}

		handleBotMessage(msg)
	}
}

func OnMatchFound(p1, p2 queueloop.Player) {
	if BotConn == nil {
		fmt.Println("No bot connected to handle match")
		return
	}

	matchMsg := map[string]any{
		"type": "match_found",
		"player1": map[string]any{
			"userId": p1.UserId,
			"elo":    p1.Elo,
		},
		"player2": map[string]any{
			"userId": p2.UserId,
			"elo":    p2.Elo,
		},
	}

	err := BotConn.WriteJSON(matchMsg)
	if err != nil {
		fmt.Println("Error sending match_found message to bot:", err)
	}
}

func handleBotMessage(msg map[string]any) {
	switch msg["type"] {
	case "queue_player":
		userId := msg["userId"].(string)
		elo := int(msg["elo"].(float64))
		var player = queueloop.Player{UserId: userId, Elo: elo, TimeJoined: time.Now()}

		queueloop.Add(player)
		fmt.Printf("Received queuePlayer: %s (%d)\n", userId, elo)
	case "dequeue_player":
		userId := msg["userId"].(string)
		queueloop.Remove(userId)
	case "bot_ready":
		fmt.Println("Bot is ready to receive matches")
	}
}
