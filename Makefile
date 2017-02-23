install:
	@go get gopkg.in/redis.v5
	@go build -o $GOPATH/bin/train github.com/huacnlee/train/cmd
	@cd app; go get
server:
	revel run github.com/fdd/website-m-backend
release:
	@make assets
	GOOS=linux GOARCH=amd64 revel package github.com/huacnlee/mediom prod
assets:
	@train --source app/assets --out public
test:
	@cd app; go test
	@cd app/controllers; go test
	@cd app/models; go test
	@cd tests; go test
