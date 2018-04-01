dependencies:
	go get github.com/golang/dep/cmd/dep
	dep ensure -v

	go get github.com/alecthomas/gometalinter
	gometalinter --install

spec:
	go test -v -tags=spec -cover ./...

feature:
	go test -v -tags=feature -cover ./...

analysis:
	gometalinter --vendor ./... --disable-all --enable=gofmt --enable=golint --enable=vet --enable=errcheck
