#

docker compose up --build -d
docker compose exec -it ms-walletcore go run cmd/walletcore/main.go
abra o client.http
    crie os clients, accounts e transactions
para visualizar o banco de dados, acesse http://localhost:8081
    adiciona mais valores no balance do account
para visualizar os eventos enviados para o kafka, acesse http://localhost:9021
    clique em "Topics"
    seleciona o topico "transactions" ou "balances"
    selecione messages
    em offset, coloque "0" e de enter