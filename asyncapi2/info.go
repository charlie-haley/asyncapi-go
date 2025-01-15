package asyncapi2

type Info struct {
	Title       string `json:"title"`
	Version     string `json:"version"`
	Description string `json:"description,omitempty"`
}

func NewInfo() *Info {
	return &Info{}
}

func (i *Info) WithTitle(title string) *Info {
	i.Title = title
	return i
}

func (i *Info) WithVersion(version string) *Info {
	i.Version = version
	return i
}

func (i *Info) WithDescription(description string) *Info {
	i.Description = description
	return i
}
