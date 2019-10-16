IMAGE_TAG=thenets/backup

dev-env: build
	docker run -it --rm \
		--name backup-dev-env \
		-v $(PWD):/app \
		--entrypoint="/bin/bash" \
		$(IMAGE_TAG)

build:
	docker build -t $(IMAGE_TAG) .
