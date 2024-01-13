test-coverage:
	go test ./... -v -coverprofile=coverage.out -run TestCaptcha

test:
	go test ./... -v -run TestCaptcha
