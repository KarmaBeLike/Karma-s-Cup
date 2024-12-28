package content

type ContentProvider interface {
	GetContent() (string, error)
}
