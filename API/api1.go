package API

type Namer interface {
	Show() string
	Cancel() bool
}
