HTTP=:6060

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	go test -cover ./...

.PHONY: doc
doc:
	go tool godoc -http=$(HTTP)
