export tag=v1

build:
	echo "building service2 container"
	docker build -t cr7258/service2:${tag} .

push: build
	echo "pushing cr7258/service2"
	docker push cr7258/service2:${tag}