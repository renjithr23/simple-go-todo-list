check_install: 
	where swagger || go get github.com/go-swagger/go-swagger/cmd/swagger

swagger: 
	swagger generate spec -o ./swagger.yaml --scan-models