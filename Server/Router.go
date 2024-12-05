package Server

type Route struct {
	Path    string
	Mapping map[string]string
	Entity  string
}
