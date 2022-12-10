package notification

type Client interface {
	IsAllowed(Notification) bool
	Register() chan Notification
	Send(Notification)
}

type FnIsAllowed func(Notification) bool

type DefaultClient struct {
	isAllowed     FnIsAllowed
	notifications chan Notification
}

func (c *DefaultClient) IsAllowed(n Notification) bool {
	if c.isAllowed != nil {
		return c.isAllowed(n)
	}

	return false
}

func (c *DefaultClient) Register() chan Notification {
	return c.notifications
}

func (c *DefaultClient) Send(n Notification) {
	c.notifications <- n
}

func NewDefaultClient(isAllowed FnIsAllowed) *DefaultClient {
	return &DefaultClient{
		isAllowed:     isAllowed,
		notifications: make(chan Notification),
	}
}
