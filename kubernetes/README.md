# Módulo - Kubernetes

## Execução com Docker

- `docker compose up -d --build`
- Acesse o `localhost:8000`

## Criando cluster, context e namespace

- `kind create cluster fullcycle`
- `kubectl config view`
- `kubectl create namespace ns-fullcycle`
- `kubectl config set-context fullcycle --namespace=ns-fullcycle --cluster=kind-fullcycle --user=kind-fullcycle`
- `kubectl config use-context fullcycle`
- `kubectl config get-contexts`

## Execução com Kubernetes

### Deployment, HPA

- `kubectl apply -f k8s/deployment.yaml -f k8s/service.yaml -f k8s/metrics-server.yaml -f k8s/hpa.yaml -f k8s/pvc.yaml -f k8s/configmap-env.yaml`
- `kubectl port-forward service/goserver 8000:8000`
- `watch -n 1 "kubectl get all"`
- `watch -n 1 "kubectl top pod"`

- `kubectl rollout history deployment goserver`: visualizar o histórico do deployment
- `kubectl rollout undo deployment goserver --to-revision=2`: voltar ao deployment de identificador `2`

#### Teste de carga com Fortio

- `kubectl run -it fortio --rm --image=fortio/fortio -- load -qps 800 -t 120s -c 70 "http://goserver:8000/healthz"`

### StatefulSet

> StatefulSet cria pods de forma ordenada e deleta em ordem inversa da de criação. Service Headless é usado para criar ponto de acesso em um serviço especifico, atraves do dns e sem loadbalancer `mysql-0.mysql-headless`.
> Volumes dinâmicos que são criados conforme novos pods vão sendo criados ou volumes sendo reatachados conforme pods são destruidos e recriados.

- `kubectl apply -f k8s/statefulset.yaml -f k8s/service-headless.yaml`
- `watch -n 1 "kubectl get all"`
- `watch -n 1 "kubectl get pvc"`

## Usar o DevContainers do VSCode

- Abre a pasta `./` no vscode
- Instale a extensão Dev Containers
- `ctlr+shift+p` `Rebuild and Reopen in Container`

network policy, deploy, afinity, rbac, resources
