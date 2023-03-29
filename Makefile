PROJECT:=go-admin

.PHONY: build
build:
	CGO_ENABLED=0 go build -o go-admin main.go
build-sqlite:
	go build -tags sqlite3,json1 -o go-admin main.go
publish-sqlite:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags sqlite3,json1 -o go-admin main.go
#.PHONY: test
#test:
#	go test -v ./... -cover

#.PHONY: docker
#docker:
#	docker build . -t go-admin:latest

gencode:
	./go-cli generate mysql -u 'sciuse:SOXTW5Kbq2CPmdRrshyS@tcp(rm-uf66ce9pg3zf971zovo.mysql.rds.aliyuncs.com:3608)/sciuseapi' -t sys_user
