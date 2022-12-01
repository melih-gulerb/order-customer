package configs

type Config struct {
	MongoConnection    string
	DBName             string
	CustomerCollection string
}

var cfg *Config

func SetConfig() {
	cfg = &Config{
		MongoConnection:    "mongodb+srv://gg-username:gg.@cluster0.w0wd4jk.mongodb.net/test",
		DBName:             "nexus",
		CustomerCollection: "customer",
	}
}

func GetConfig() *Config {
	return cfg
}
