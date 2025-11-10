package config

var config = make(map[string]string)

func Init(host, port, authToken string) {
	config["token"] = authToken
	config["host"] = host
	config["port"] = port
}

func Host() string {
	return config["host"]
}

func Port() string {
	return config["port"]
}

func Token() string {
	return config["token"]
}
