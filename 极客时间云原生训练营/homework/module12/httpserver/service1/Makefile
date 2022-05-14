export tag=v1

build:
	echo "building service1 container"
	docker build -t cr7258/service1:${tag} .

push: build
	echo "pushing cr7258/service1"
	docker push cr7258/service1:${tag}