# README

- Crie a pasta do volume para não dar erro de permissão
    - `mkdir elasticsearch_data`
- Crie uma network que será usada
    - `docker network create --driver bridge observability`
- Inicie todos os serviços
    - `docker compose up -d`