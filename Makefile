export GOOS=linux
export GOARCH=amd64
export IMAGE_PREFIX=vinkdong
.PHONY: directories
SEDMID := 

directories: ${OUT_DIR}

${OUT_DIR}:
	mkdir -p build/bin

server: directories
	go build -o build/bin/server cmd/server/server.go
	
dockerfiledir:
	rm -rf build/tmp/${APP}
	mkdir -p build/tmp/${APP}
	
dockerfile: dockerfiledir
	cp Dockerfile.tpl build/tmp/${APP}/Dockerfile
	cp -r build/bin/${APP} build/tmp/${APP}
	sed -i 'bk' "s/{{APP_NAME}}/${APP}/g" build/tmp/${APP}/Dockerfile
	
image: server dockerfile
	docker build -t ${IMAGE_PREFIX}/${APP} build/tmp/${APP}/

push: image
	docker push ${IMAGE_PREFIX}/${APP}