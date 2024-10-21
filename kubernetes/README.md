# Módulo - Kubernetes

## Execução com Docker

- `docker compose up -d --build`
- Acesse o `localhost:8000`

## Execução com Kubernetes

- `kind create cluster --config ./k8s/kind.yaml --name cluster-fullcycle`
- `kubectl create namespace ns-fullcycle`
- `kubectl config set-cluster cluster-fullcycle`
- `kubectl config use-context ns-fullcycle`
- `kubectl apply -f k8s/pod.yaml`
- `kubectl port-forward pod/goserver 8000:8000`

### Deployment commands

- `kubectl rollout history deployment goserver`
- `kubectl rollout undo deployment goserver --to-revision=2`

## Usar o DevContainers do VSCode

- Abre a pasta `./` no vscode
- Instale a extensão Dev Containers
- `ctlr+shift+p` `Rebuild and Reopen in Container`
