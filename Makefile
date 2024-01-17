test-coverage:
	go test ./... -v -coverprofile=coverage.out

test:
	go test ./... -v

test-captcha:
	go test ./... -v -run TestCaptcha
