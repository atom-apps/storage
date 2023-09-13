//go:generate atomctl gen routes
//go:generate swag fmt
//go:generate swag init -ot json
package main

import (
	"log"

	"github.com/atom-apps/storage/database/query"
	"github.com/atom-apps/storage/modules/boot"
	moduleStorage "github.com/atom-apps/storage/modules/storages"
	database "github.com/atom-providers/database-mysql"
	"github.com/atom-providers/jwt"
	serviceHttp "github.com/atom-providers/service-http"
	"github.com/atom-providers/uuid"

	"github.com/rogeecn/atom"
)

func main() {
	providers := serviceHttp.
		Default(
			uuid.DefaultProvider(),
			jwt.DefaultProvider(),
			query.DefaultProvider(),
			database.DefaultProvider(),
		).
		With(
			boot.Providers(),
			moduleStorage.Providers(),
		)

	opts := []atom.Option{
		atom.Name("storage"),
		atom.RunE(serviceHttp.ServeE),
	}

	if err := atom.Serve(providers, opts...); err != nil {
		log.Fatal(err)
	}
}
