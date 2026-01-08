package asyncapi3

type Reference struct {
	Ref string `json:"$ref,omitempty"`
}

func NewReference(ref string) *Reference {
	return &Reference{Ref: ref}
}

func ChannelRef(channelName string) *Reference {
	return NewReference("#/channels/" + channelName)
}

func ServerRef(serverName string) *Reference {
	return NewReference("#/servers/" + serverName)
}

func MessageRef(channelName, messageName string) *Reference {
	return NewReference("#/channels/" + channelName + "/messages/" + messageName)
}

func ComponentMessageRef(messageName string) *Reference {
	return NewReference("#/components/messages/" + messageName)
}

func ComponentSchemaRef(schemaName string) *Reference {
	return NewReference("#/components/schemas/" + schemaName)
}
