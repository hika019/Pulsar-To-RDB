package config

type Config struct {
	Input  Input  `json:"input"`
	Output Output `json:"output"`
}

type Input struct {
	Host  string `json:"host"`
	Topic string `json:"topic"`
	Codec string `json:"codec"`
}

type Output struct {
	File File `json:"file"`
	Rdb  Rdb  `json:"RDB"`
}

type File struct {
	Path     string `json:"path"`
	Filename string `json:"filename"`
}

type Rdb struct {
	Host      string   `json:"host"`
	Schema    string   `json:"schema"`
	User      string   `json:"user"`
	Password  string   `json:"password"`
	Statement []string `json:"statement"`
}

type Env struct {
	ConfPath string
}
