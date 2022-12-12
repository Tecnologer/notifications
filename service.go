package notification

import "sync"

//Service is a struct to manage a service of notification
type Service struct {
	rw                    sync.RWMutex
	isRunning             bool
	incomingNotifications chan Notification
	clients               []Client
	wg                    sync.WaitGroup
}

//NewService creates a new instance of service
func NewService() *Service {
	return &Service{
		incomingNotifications: make(chan Notification, 1),
		clients:               make([]Client, 0),
	}
}

//RegisterClients registers the clients to send the notifications
func (s *Service) RegisterClients(newClients ...Client) {
	s.rw.Lock()
	defer s.rw.Unlock()

	s.wg.Add(len(newClients))
	s.clients = append(s.clients, newClients...)

	//start the service on the firts client registered
	if !s.isRunning {
		go s.distribute()
	}
}

//Send sends a notification to the service
func (s *Service) Send(notification Notification) {
	s.incomingNotifications <- notification
}

//CloseNWait closes the channel to not receive more notifications and wait until all notification is send to each client
func (s *Service) CloseNWait() {
	s.isRunning = false
	close(s.incomingNotifications)
	s.wg.Wait()
}

//disitribute is a function that reads the incoming notification and distribute it to each client
func (s *Service) distribute() {
	s.isRunning = true
	for notif := range s.incomingNotifications {
		s.rw.RLock()
		for _, client := range s.clients {
			if client.IsClosed() || !client.IsAllowed(notif) {
				continue
			}

			client.Send(notif)
		}
		s.rw.RUnlock()
	}

	s.closeClients()
}

//cloaseClients calls the close method of each client
func (s *Service) closeClients() {
	for _, client := range s.clients {
		s.wg.Done()
		client.Close()
	}
}
