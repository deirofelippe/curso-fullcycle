go mod tidy
go test ./...
go get
go init github.com/deirofelippe/curso-fullcycle

docker compose exec mysql mysql -uroot -p wallet -h localhost -e \
    "CREATE TABLE IF NOT EXISTS clients(id varchar(255), name varchar(255), email varchar(255), created_at date) ;
    CREATE TABLE IF NOT EXISTS accounts(id varchar(255), client_id varchar(255), balance int, created_at date) ;
    CREATE TABLE IF NOT EXISTS transactions(id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date)"

docker compose exec mysql mysql -uroot -p wallet -e \
    "USE wallet ; SHOW tables"

docker compose exec mysql mysql -uroot -p wallet

docker compose exec -it kafka bash -c \
    "kafka-topics --bootstrap-server localhost:9092 --topic balances --create ;
    kafka-topics --bootstrap-server localhost:9092 --topic transactions --create"

go run cmd/walletcore/main.go
