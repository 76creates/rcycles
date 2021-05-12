package resource

type Resource interface{
	Lookup() error
}
