package config

type Config struct {
	Input  Pulsar `yaml:"input"`
	Output Output `yaml:"output"`
}

type Pulsar struct {
	Host  string `yaml:"host"`
	Topic string `yaml:"topic"`
	Codec string `yaml:"codec"`
}

type Output struct {
	File File `yaml:"file"`
	Rdb  Rdb  `yaml:"rdb"`
}

type File struct {
	Path     string `yaml:"path"`
	Filename string `yaml:"filename"`
}

type Rdb struct {
	Driver    string   `yaml:"driver"`
	Host      string   `yaml:"host"`
	Schema    string   `yaml:"schema"`
	User      string   `yaml:"user"`
	Password  string   `yaml:"password"`
	Statement []string `yaml:"statement"`
}

type Env struct {
	ConfPath string
}
