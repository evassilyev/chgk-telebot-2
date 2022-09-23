package core

type Configuration struct {
	Domain   string `default:"http://localhost:8080"` // Domain name in telegram webhook
	Url      string `default:"/webhook"`              // Webhook url
	CertFile string `required:"true" split_words:"true"`
	BotToken string `required:"true" split_words:"true"`
	Port     uint16 `default:"8080"` // Listened port
	KeyFile  string `required:"true" split_words:"true"`
	DbUser   string `required:"true" split_words:"true"`
	DbPass   string `required:"true" split_words:"true"`
	DbName   string `required:"true" split_words:"true"`
	DbPort   uint16 `required:"true" split_words:"true"`
	DbHost   string `required:"true" split_words:"true"`
}
