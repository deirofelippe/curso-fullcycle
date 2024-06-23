# README

- Crie a pasta do volume para não dar erro de permissão
    - `mkdir elasticsearch_data`
- Crie uma network que será usada
    - `docker network create --driver bridge observability`
- Inicie todos os serviços
    - `docker compose up -d`
- O arquivo `metricbeat.yml` deve ter as permissões: - rw- r-- r-- root:root
    - sudo chown root:root metricbeat.yml
    - sudo chmod 644 metricbeat.yml
- O arquivo `heartbeat.yml` deve ter as permissões: - rw- r-- r-- 
    - sudo chmod 644 heartbeat.yml
