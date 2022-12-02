package configs

type Config struct {
	MongoConnection    string
	DBName             string
	CustomerCollection string
	OrderCollection    string
	OrderClient        string
	CustomerClient     string
}

var cfg *Config

func SetConfig() *Config {
	cfg = &Config{
		MongoConnection:    "mongodb+srv://gg-username:gg.@cluster0.w0wd4jk.mongodb.net/test",
		DBName:             "nexus",
		CustomerCollection: "customer",
		OrderCollection:    "order",
		CustomerClient:     "http://localhost:4000",
		OrderClient:        "http://localhost:2000",
	}

	return cfg
}
