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

## RBAC

- kubectl apply -f k8s/deployment.yaml -f k8s/security/goserver-account.yaml

- Cria o context com namespace e cluster
- kubectl config set-credentials user-admin --client-certificate=user-admin.crt --client-key=user-admin.key
- kubectl config set-credentials user-read --client-certificate=user-read.crt --client-key=user-read.key

- kubectl config set-context admin-context --cluster=kind-fullcycle --namespace=ns-fullcycle --user=user-admin
- kubectl config set-context read-context --cluster=kind-fullcycle --namespace=ns-fullcycle --user=user-read

- kubectl config use-context admin-context

- Criar User
  - openssl genrsa -out admin-user.key 2048
  - openssl req -new -key admin-user.key -out admin-user.csr -subj "/CN=admin-user"
  - Execute o comando para criar o CSR

```bash
cat <<EOF | kubectl apply -f -
apiVersion: certificates.k8s.io/v1
kind: CertificateSigningRequest
metadata:
    name: admin-user-csr
spec:
    request: $(cat admin-user.csr | base64 | tr -d '\n')
    signerName: kubernetes.io/kube-apiserver-client
    usages:
    - client auth
EOF
```

    - kubectl certificate approve admin-user-csr
    - kubectl get csr admin-user-csr -o jsonpath='{.status.certificate}' | base64 -d > admin-user.crt

- Criar ServiceAccount

kubectl config set-credentials admin-user --client-certificate=admin-user.crt --client-key=admin-user.key
kubectl config set-context admin-fullcycle --cluster=kind-fullcycle --user=admin-user --namespace=ns-fullcycle
kubectl config use-context admin-fullcycle

## Usar o DevContainers do VSCode

- Abre a pasta `./` no vscode
- Instale a extensão Dev Containers
- `ctlr+shift+p` `Rebuild and Reopen in Container`

network policy, deploy, afinity, rbac, resources
