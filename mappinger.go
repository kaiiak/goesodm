package goesodm

// Mappinger is a elastic search mapping
type Mappinger interface {
	Mapping() string
	Index() string
	Type() string
	Init() Initer
}

// Aliaser is a single alias of elastic search
type Aliaser interface {
	Mappinger
	Alias() string
}

// Initer with
type Initer interface {
	Write2Es() error
	Next() bool
}
