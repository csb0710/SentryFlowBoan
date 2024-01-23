IMAGE_NAME = boanlab.com/custom-collector
TAG = v1

.PHONY: build

build:
	docker build -t $(IMAGE_NAME):$(TAG) .

.PHONY: clean

clean:
	docker rmi $(IMAGE_NAME):$(TAG)
