package Camera

type ConnectionStatus string

const (
	ConnectionStatusNotConnectedYet = "not_connected_yet"
	ConnectionStatusOnline          = "online"
	ConnectionStatusOffline         = "offline"
	ConnectionStatusInvalidAuth     = "invalid_auth"
)

type Status string

const (
	StatusActive   Status = "active"
	StatusInActive Status = "inactive"
)

type Camera struct {
	Uuid                 string
	Title                string
	BaseUrl              string
	User                 string /* Camera user login */
	Password             string /* Camera password */
	LastConnectionStatus ConnectionStatus
	Status               Status
}
