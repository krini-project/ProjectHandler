GOPATH=~/go

~/go/bin/swag init -g api/server/Server.go api/server/CommonHandler.go models/Models.go
cp docs/swagger.json www/swagger-ui