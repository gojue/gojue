package gojue

type goJueConfig struct {
	Kernelmajor byte
	Kernelminor byte
	Kernelpatch byte
	btfEnable   bool
	enableCORE  bool
}
