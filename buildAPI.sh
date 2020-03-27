GOPATH=~/go

~/go/bin/swag init -g server/Server.go server/CommonHandler.go
cp docs/swagger.json www/swagger-ui