export tag=v1

build:
	echo "building service0 container"
	docker build -t cr7258/service0:${tag} .

push: build
	echo "pushing cr7258/service0"
	docker push cr7258/service0:${tag}