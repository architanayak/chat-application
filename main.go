package main

import (
	"fmt"
	"log"
	"net/http"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	router *mux.Router
	Server *gosocketio.Server
)

type Message struct {
	Text string `json:"message"`
}

func init() {
	Server = gosocketio.NewServer(transport.GetDefaultWebsocketTransport())
	fmt.Println("Socket Inititalize...")
}

func LoadSocket() {
	// socket connection
	Server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		fmt.Println("Connected", c.Id())
		c.Join("Room")
	})

	// socket disconnection
	Server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		fmt.Println("Disconnected", c.Id())

		// handles when someone closes the tab
		c.Leave("Room")
	})

	// chat socket
	Server.On("/chat", func(c *gosocketio.Channel, message Message) string {
		fmt.Println(message.Text)
		c.BroadcastTo("Room", "/message", message.Text)
		return "message sent successfully."
	})
}

func CreateRouter() {
	router = mux.NewRouter()
}

func InititalizeRoutes() {
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	})
	router.Handle("/socket.io/", Server)
}

func StartServer() {
	fmt.Println("Server Started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}

func main() {
	LoadSocket()
	CreateRouter()
	InititalizeRoutes()
	StartServer()
}
