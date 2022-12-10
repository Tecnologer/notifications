package main

import (
	"fmt"

	"github.com/tecnologer/notification"
)

type client struct {
	*notification.DefaultClient
	name string
}

func newClient(name string, fn notification.FnIsAllowed) *client {
	return &client{
		DefaultClient: notification.NewDefaultClient(fn),
		name:          name,
	}
}

func main() {
	notif := []notification.Notification{
		notification.NewDefault("notification 1", "type 1"),
		notification.NewDefault("notification 2", "type 4"),
		notification.NewDefault("notification 3", "type 2"),
		notification.NewDefault("notification 4", "type 1"),
		notification.NewDefault("notification 5", "type 3"),
	}

	client1 := newClient("client 1", isAllowedClient1)                                       //only type 1
	client2 := newClient("client 2", isAllowedClient2)                                       //only type 2
	client3 := newClient("client 3", func(n notification.Notification) bool { return true }) //all notifications
	client4 := newClient("client 4", nil)                                                    //none notification

	notification.RegisterClient(client1, client2, client3, client4)

	register(client1)
	register(client2)
	register(client3)
	register(client4)

	for _, n := range notif {
		notification.Send(n)
	}

	fmt.Scanln()
}

func isAllowedClient1(n notification.Notification) bool {
	return n.GetType() == "type 1"
}

func isAllowedClient2(n notification.Notification) bool {
	return n.GetType() == "type 2"
}

func register(c notification.Client) {
	go func() {
		for n := range c.Register() {
			client := c.(*client)
			fmt.Printf("%s: %s, %s\n", client.name, n.Get(), n.GetType())
		}
	}()
}
