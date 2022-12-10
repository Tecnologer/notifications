package notification

import "sync"

var (
	rw                    sync.RWMutex
	isRunning             bool
	incomingNotifications = make(chan Notification, 1)
	clients               = make([]Client, 0)
)

type Notification interface {
	Get() string
	GetType() string
}

type Default struct {
	Msg  string
	Type string
}

func NewDefault(msg, t string) *Default {
	return &Default{
		Msg:  msg,
		Type: t,
	}
}

func (dn *Default) Get() string {
	return dn.Msg
}

func (dn *Default) GetType() string {
	return dn.Type
}

func RegisterClient(client ...Client) {
	rw.Lock()
	defer rw.Unlock()

	clients = append(clients, client...)

	//start the service on the firts client registered
	if !isRunning {
		go distribute()
	}
}

func Send(notification Notification) {
	incomingNotifications <- notification
}

func distribute() {
	isRunning = true
	for notif := range incomingNotifications {
		rw.RLock()
		for _, client := range clients {
			if !client.IsAllowed(notif) {
				continue
			}

			client.Send(notif)
		}
		rw.RUnlock()
	}
}
