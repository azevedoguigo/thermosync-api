package handler

import (
	"net/http"

	"github.com/azevedoguigo/thermosync-api/internal/websocket"
)

func Websocket(w http.ResponseWriter, r *http.Request) {
	websocket.HanldeConnections(w, r)
}
