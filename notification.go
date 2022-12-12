//notification is a package to define a service to manage and send notification to the clients
//that are register for that
package notification

//Notification is a interface to get the body of notification
//and a type to identify it
type Notification interface {
	Get() string
	GetType() string
}

//Default is a strutc for a basic notification
type Default struct {
	Msg  string
	Type string
}

//NewDefault creates a new instance of a Default Notification
func NewDefault(msg, t string) *Default {
	return &Default{
		Msg:  msg,
		Type: t,
	}
}

//Get returns the body of the notification
func (dn *Default) Get() string {
	return dn.Msg
}

//GetType returns the type of the notification
func (dn *Default) GetType() string {
	return dn.Type
}
