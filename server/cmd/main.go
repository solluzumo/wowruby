package main

import (
	"database/sql"
	applicationService "servre/internal/application/services"
	domainRepository "servre/internal/domain/repository"
)

func main() {
	dbType := "postgres"

	var itemRepo domainRepository.ItemRepository

	switch dbType {
	case "postgres":
		connStr := "user=postgres password=mypass dbname=productdb sslmode=disable"
		db, err := sql.Open("postgres", connStr)

	}

	itemApp := applicationService.NewItemService(itemRepo)

	go rest.StartServer(itemApp)
}
