package walletcore

import (
	"database/sql"
	"fmt"

	"github.com/deirofelippe/curso-fullcycle/internal/database"
	"github.com/deirofelippe/curso-fullcycle/internal/event"
	createaccount "github.com/deirofelippe/curso-fullcycle/internal/usecase/create_account"
	createclient "github.com/deirofelippe/curso-fullcycle/internal/usecase/create_client"
	createtransaction "github.com/deirofelippe/curso-fullcycle/internal/usecase/create_transaction"
	"github.com/deirofelippe/curso-fullcycle/pkg/events"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost", "3306", "wallet"))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()

	eventDispatcher.Register("TransactionCreated", handler)

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)
	transactionDb := database.NewTransactionDB(db)

	createClientUsecase := createclient.NewCreateClientUsecase(clientDb)
	createAccountUsecase := createaccount.NewCreateAccountUsecase(accountDb, clientDb)
	createTransactionsUsecase := createtransaction.NewCreateTransactionUsecase(transactionDb, accountDb, eventDispatcher, transactionCreatedEvent)
}
