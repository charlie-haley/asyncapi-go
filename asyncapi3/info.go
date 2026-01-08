package asyncapi3

type Info struct {
	Title          string   `json:"title"`
	Version        string   `json:"version"`
	Description    string   `json:"description,omitempty"`
	TermsOfService string   `json:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty"`
	License        *License `json:"license,omitempty"`
	Tags           []Tag    `json:"tags,omitempty"`
}

type Contact struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

type License struct {
	Name string `json:"name"`
	URL  string `json:"url,omitempty"`
}

func NewInfo() *Info {
	return &Info{
		Tags: make([]Tag, 0),
	}
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

func (i *Info) WithContact(contact *Contact) *Info {
	i.Contact = contact
	return i
}

func (i *Info) WithLicense(license *License) *Info {
	i.License = license
	return i
}

func (i *Info) WithTag(tag Tag) *Info {
	i.Tags = append(i.Tags, tag)
	return i
}
