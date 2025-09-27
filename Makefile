run:
	go run main.go

cluster:
	minikube start -p ralvescosta --nodes 2 --kubernetes-version=stable --driver=docker