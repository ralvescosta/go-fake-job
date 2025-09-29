run:
	go run main.go

cluster:
	minikube start -p ralvescosta --nodes 2 --kubernetes-version=stable --driver=docker

build:
	docker build --network=host . -t rafaelbodao/fakejob:latest

push:
	docker push rafaelbodao/fakejob:latest