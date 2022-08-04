package api

// SSL specification app
type SSL struct {
	Enable bool   `required:"true"`
	Cert   string `required:"false"`
	Key    string `required:"false"`
}

// Config support to the server
type Config struct {
	Ssl  SSL    `required:"true"`
	Port string `required:"true"`
}
