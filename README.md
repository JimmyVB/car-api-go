# Requirements

- go v1.17
- go modules
- swag v1.7.2

# Build

- Install dependencies:
  `$ go mod init car-api`

- Clean dependencies:
  `$ go mod tidy`

- Install swag
  `$ go get -u github.com/swaggo/swag/cmd/swag@v1.7.2`

- Generate swagger docs
  `$ swag init`

- Run app :
  `$ go run main.go`

- Build
  `$ go build -o car-api main.go`

# Environments

#### required environment variables

- `PORT`: port for the server
- `DB_HOST:` host of postgresql
- `DB_PORT:` port of postgresql
- `DB_NAME:` database name
- `DB_USER:` user of the database
- `DB_PASSWORD:` user password

# Docker

####Create image
`$ docker build -t car-api .`

####Create contanier
`$ docker run -p 8080:8080 --name car-api car-api`

#Kubernetes
The deployment in kubernetes is through minikube to simulate a local server. Check the official page to install: https://minikube.sigs.k8s.io/docs/start/

###Start minikube
`$ start minikube`

###Deploy
####Deploy yaml
`$ kubectl create -f deployment.yaml`

####Get deployments
`$ kubectl get deployments`

####Get pods
`$ kubectl get pods`

####Get pods
`$ kubectl get pods`

###Expose
####Expose deployment
`$ kubectl expose deployment acl-car-api --type=NodePort --name=acl-car-api --target-port=8080 `

####Expose url
`$ minikube service acl-car-api --url`
