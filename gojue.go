package gojue

type goJueSrv struct {
	gjConfig *goJueConfig
	modules  map[string]IModule
}

type (
	// RunOption defines the method to customize a Server.
	RunOption func(*goJueSrv)
)

func NewJueServer(c *goJueConfig, opts ...RunOption) (*goJueSrv, error) {
	server := &goJueSrv{
		gjConfig: c,
	}

	//opts = append([]RunOption{WithNotFoundHandler(nil)}, opts...)
	//for _, opt := range opts {
	//	opt(server)
	//}
	return server, nil
}

/*
// AddModules add given modules into the Server.
func (s *jueSrv) AddModules(rs []Route, opts ...ModuleOption) {
	...
}
*/
