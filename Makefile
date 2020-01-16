mock:
	go install github.com/golang/mock/mockgen
	mockgen -destination mock_go_mgt_client/mock_mgt_client.go github.com/Joingo/go-mgt-client Client

.PHONY: test
test:
	go test ./...
