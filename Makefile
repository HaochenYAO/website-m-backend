install:
	@go get gopkg.in/redis.v5
	@go get github.com/go-xorm/xorm
server:
	revel run github.com/website-m-backend
test:
	@cd app; go test
	@cd app/controllers; go test
	@cd app/models; go test
	@cd tests; go test
