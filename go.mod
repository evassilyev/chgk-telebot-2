module github.com/evassilyev/chgk-telebot-2

go 1.19

require (
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1
	github.com/gorilla/mux v1.8.0
	github.com/jmoiron/sqlx v1.3.5
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/lib/pq v1.10.7
)

require (
	github.com/google/uuid v1.3.0
	github.com/kenshaw/inflector v0.2.0
	github.com/kenshaw/snaker v0.2.0
	github.com/stretchr/testify v1.7.0
	github.com/xo/xo v0.0.0-00010101000000-000000000000
	golang.org/x/tools v0.2.0
	mvdan.cc/gofumpt v0.4.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/cobra v1.6.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/mod v0.6.0 // indirect
	golang.org/x/sys v0.1.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/xo/xo => github.com/evassilyev/xo v1.1.1
