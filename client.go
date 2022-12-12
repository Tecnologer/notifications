package notification

type Client interface {
	//IsAllowed is a function to verify is the client can see the notification
	IsAllowed(Notification) bool
	//Register registers the client as candidate to receive candidate
	Register() chan Notification
	//Send sends a notification to the client
	Send(Notification)
	//Close unregister the client for recieve notifications
	Close()
	//IsClosed returns true if the client is not available for receive notifications
	IsClosed() bool
}

//FnIsAllowed is a function to type to validate is the client can receive the notification
type FnIsAllowed func(Notification) bool

//DefaultClient is a struct with a basic client's information
type DefaultClient struct {
	isClosed      bool
	isAllowed     FnIsAllowed
	notifications chan Notification
}

//NewDefaultClient creates a new instance of DefaultClient
func NewDefaultClient(isAllowed FnIsAllowed) *DefaultClient {
	return &DefaultClient{
		isAllowed:     isAllowed,
		notifications: make(chan Notification),
	}
}

//IsAllowed execute the FnIsAllowed, if the function is not defined returns false
func (c *DefaultClient) IsAllowed(n Notification) bool {
	if c.isAllowed != nil {
		return c.isAllowed(n)
	}

	return false
}

//Register returns a channel to receive the notifications for this client
func (c *DefaultClient) Register() chan Notification {
	return c.notifications
}

//Send sends a notification for this client
func (c *DefaultClient) Send(n Notification) {
	c.notifications <- n
}

//Close marks this client not available to recieve more notifications
func (c *DefaultClient) Close() {
	close(c.notifications)
	c.isClosed = true
}

//IsClosed returns if the client is available to recive notifications
func (c *DefaultClient) IsClosed() bool {
	return c.isClosed
}
