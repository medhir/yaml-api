.PHONY: app
app: 
	go run app.go

.PHONY: test
test: 
	go test ./...