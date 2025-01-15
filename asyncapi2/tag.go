package asyncapi2

type Tag struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

func NewTag(name string) *Tag {
	return &Tag{Name: name}
}

func (t *Tag) WithDescription(description string) *Tag {
	t.Description = description
	return t
}
