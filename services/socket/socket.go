package socket

import (
	"bifrost/models"
	"bifrost/services/db"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/rs/cors"
	socketio "github.com/vchitai/go-socket.io/v4"
	"github.com/vchitai/go-socket.io/v4/engineio"
	"github.com/vchitai/go-socket.io/v4/engineio/transport"
	"github.com/vchitai/go-socket.io/v4/engineio/transport/polling"
	"github.com/vchitai/go-socket.io/v4/engineio/transport/websocket"
	"gorm.io/gorm"
)

var Server *socketio.Server
var userConnections = make(map[string]socketio.Conn)

type SocketRepository interface {
	BroadcastToRoom(room string, event string, message string) error
	SendMessageToUser(userID string, event string, message string) error
}

type SocketRepositoryImpl struct {
	DB *gorm.DB
}

var allowOriginFunc = func(r *http.Request) bool {
	return true
}

func ListenServer() {

	Server = socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})

	Server.OnConnect("/", func(c socketio.Conn, m map[string]interface{}) error {
		log.Println("connected:", c.ID())
		c.Join("chat")
		return nil
	})

	Server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		log.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	Server.OnDisconnect("/", func(c socketio.Conn, reason string, m map[string]interface{}) {
		fmt.Println("Disconnected:", c.ID())
	})

	go func() {
		if err := Server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer Server.Close()

	mux := http.NewServeMux()

	mux.Handle("/socket.io/", Server)
	mux.HandleFunc("/ass", serveHTML)

	handler := cors.Default().Handler(mux)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler = c.Handler(handler)
	log.Fatal(http.ListenAndServe(os.Getenv("SOCKET_PORT"), handler))

}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/index.html")
}

func (repo *SocketRepositoryImpl) BroadcastToRoom(namespace string, room string, event string, msg string) error {
	Server.BroadcastToRoom(namespace, room, event, msg)
	return nil
}

func (repo *SocketRepositoryImpl) SendMessageToUser(userID uuid.UUID, event string, message string) error {
	userRepo := &db.UserRepositoryImpl{DB: repo.DB}
	user, err := userRepo.GetUser(&models.User{ID: userID})
	if err != nil {
		return errors.New("User not found")
	}
	if conn, ok := userConnections[*user.SocketID]; ok {
		conn.Emit(event, message)
		return nil
	}
	return nil
}
