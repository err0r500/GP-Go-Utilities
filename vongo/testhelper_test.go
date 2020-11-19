package vongo

func setupDefConnection() {
	config := &Config{
		URI:      "mongodb://root:12345@localhost:27017",
		Database: "test",
		Monitor:  nil,
	}
	DBConn = &Connection{
		Config: config,
	}
}
