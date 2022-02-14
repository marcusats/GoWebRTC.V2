package server

import(
	"flag"
	"os"
	"time"
)


var (
	addr = flag.String("addr,":"", os.Getenv("PORT"), "")
	cert= flag.String("cert", "", "")
	key = flag.String("key", "", "")

)

func Run() error {
	flag.Parse()

	if *addr == ":"{
		*addr = "8000"
	}

	app.Get("/",handlers.Welcome)
	app.Get("/room/create", handlers.RoomCreate)
	app.Get("/room/:uuid", handlers.Room)
	app.Get("/room/:uuid/websocket",)
	app.Get("/room/:uuid/chat", handlers.RoomChat)
	app.Get("/room/:uuid/chat/websocket", websocket.New(handlres.RoomChatWebsocket))
	app.Get("/room/:uuid/viewer/websocket", websocket.New(handlres.RoomViewerWebsocket))
	app.Get("/stream/:ssuid", handlers.Stream)
	app.Get("/stream/:ssuid/websocket",)
	app.Get("/stream/:ssuid/chat/websocket",)
	app.Get("/stream/:ssuid/viewer/websocket",)
}