package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/deirofelippe/curso-fullcycle/internal/database"
	"github.com/deirofelippe/curso-fullcycle/internal/event"
	createaccount "github.com/deirofelippe/curso-fullcycle/internal/usecase/create_account"
	createclient "github.com/deirofelippe/curso-fullcycle/internal/usecase/create_client"
	createtransaction "github.com/deirofelippe/curso-fullcycle/internal/usecase/create_transaction"
	"github.com/deirofelippe/curso-fullcycle/internal/web"
	"github.com/deirofelippe/curso-fullcycle/internal/web/webserver"
	"github.com/deirofelippe/curso-fullcycle/pkg/events"
	"github.com/deirofelippe/curso-fullcycle/pkg/uow"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost", "3306", "wallet"))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()

	// eventDispatcher.Register("TransactionCreated", handler)

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})

	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	createClientUsecase := createclient.NewCreateClientUsecase(clientDb)
	createAccountUsecase := createaccount.NewCreateAccountUsecase(accountDb, clientDb)
	createTransactionUsecase := createtransaction.NewCreateTransactionUsecase(uow, eventDispatcher, transactionCreatedEvent)

	webserver := webserver.NewWebServer(":3000")

	clientHandler := web.NewWebClientHandler(*createClientUsecase)
	accountHandler := web.NewWebAccountHandler(*createAccountUsecase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUsecase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	webserver.Start()
}
