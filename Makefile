test:
	cd crypto && go vet ./... && go test -v ./... -race -cover
	cd slice && go vet ./... && go test -v ./... -race -cover
	cd auth && go vet ./... && go test -v ./... -race -cover
