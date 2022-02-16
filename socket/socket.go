package socket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
)

var ocClient *corev1client.CoreV1Client
var namespace string

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func SocketIO(client *corev1client.CoreV1Client, namespace string) {
	fmt.Println("Go Websocket")
	setupRoutes()
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func reader(conn *websocket.Conn) {

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(p))
		err1 := conn.WriteMessage(messageType, p)
		if err1 != nil {
			log.Println(err)
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Websocket connected successfully")
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/ws", wsEndpoint)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
