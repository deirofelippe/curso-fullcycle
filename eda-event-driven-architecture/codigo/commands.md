go mod tidy
go test ./...
go get
go init github.com/deirofelippe/curso-fullcycle

docker compose exec mysql mysql -uroot -p wallet -e "CREATE TABLE clients(id varchar(255), name varchar(255), email varchar(255), created_at date);CREATE TABLE accounts(id varchar(255), client_id varchar(255), balance int, created_at date);CREATE TABLE transactions(id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date);"

docker compose exec mysql mysql -uroot -p wallet

go run cmd/walletcore/main.go
