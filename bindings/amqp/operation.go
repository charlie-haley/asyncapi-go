package amqp

// OperationBinding represents the AMQP 0-9-1 Operation Binding object.
//
// This object contains information about the operation representation in AMQP.
// +binding
type OperationBinding struct {
	Expiration     int      `json:"expiration,omitempty"`
	UserID         string   `json:"userId,omitempty"`
	CC             []string `json:"cc,omitempty"`
	Priority       int      `json:"priority,omitempty"`
	DeliveryMode   int      `json:"deliveryMode,omitempty"`
	Mandatory      bool     `json:"mandatory,omitempty"`
	BCC            []string `json:"bcc,omitempty"`
	Timestamp      bool     `json:"timestamp,omitempty"`
	Ack            bool     `json:"ack,omitempty"`
	BindingVersion string   `json:"bindingVersion,omitempty"`
}

// OperationDeliveryMode represents the different types of delivery modes.
const (
	OperationDeliveryModeTransient  = 1
	OperationDeliveryModePersistent = 2
)
