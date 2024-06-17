package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/deirofelippe/curso-fullcycle/internal/database"
	"github.com/deirofelippe/curso-fullcycle/internal/event"
	"github.com/deirofelippe/curso-fullcycle/internal/event/handler"
	createaccount "github.com/deirofelippe/curso-fullcycle/internal/usecase/create_account"
	createclient "github.com/deirofelippe/curso-fullcycle/internal/usecase/create_client"
	createtransaction "github.com/deirofelippe/curso-fullcycle/internal/usecase/create_transaction"
	"github.com/deirofelippe/curso-fullcycle/internal/web"
	"github.com/deirofelippe/curso-fullcycle/internal/web/webserver"
	"github.com/deirofelippe/curso-fullcycle/pkg/events"
	"github.com/deirofelippe/curso-fullcycle/pkg/kafka"
	"github.com/deirofelippe/curso-fullcycle/pkg/uow"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}
	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	eventDispatcher := events.NewEventDispatcher()
	handler := handler.NewTransactionCreatedKafkaHandler(kafkaProducer)
	eventDispatcher.Register("TransactionCreated", handler)

	transactionCreatedEvent := event.NewTransactionCreated()

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

	port := 8080
	webserver := webserver.NewWebServer(fmt.Sprintf(":%d", port))

	clientHandler := web.NewWebClientHandler(*createClientUsecase)
	accountHandler := web.NewWebAccountHandler(*createAccountUsecase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUsecase)

	webserver.AddHandler("/", func(res http.ResponseWriter, req *http.Request) {
		hello := struct {
			Message string
		}{
			Message: "Hello World!",
		}
		data, _ := json.Marshal(hello)

		res.Header().Set("Content-Type", "application/json")
		res.Write(data)
	})
	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Printf("Servidor criado na porta %d", port)
	webserver.Start()
}
