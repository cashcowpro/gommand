dependencies:
	go get github.com/golang/dep/cmd/dep
	go get github.com/DATA-DOG/godog/cmd/godog
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/onsi/gomega/...
	dep ensure -v

spec:
	go test -v -tags=spec -cover ./...

feature:
	go test -v -tags=feature -cover ./...
