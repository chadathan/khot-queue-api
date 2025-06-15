package models

type Config struct {
	Port string
	Env  string
}

type JsonConfig struct {
	DbConnection string `json:"DbConnection"`
}
