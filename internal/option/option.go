package option

// Option is used to pass optional arguments to
// the ZapRotateLogger constructor
type Option interface {
	Name() string
	Value() interface{}
}

type _option struct {
	name  string
	value interface{}
}

// New helps to declare new Option parameter
func New(name string, value interface{}) Option {
	return &_option{
		name:  name,
		value: value,
	}
}

func (o *_option) Name() string {
	return o.name
}

func (o *_option) Value() interface{} {
	return o.value
}
