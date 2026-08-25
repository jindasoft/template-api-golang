init:
	@echo "Initializing..."
	@pre-commit install
	@go mod tidy

dev:
	@echo "Running..."
	@make swag
	@air

## Test the application
test:
	@echo "Testing..."
	@go test ./... -v

## Test cover
test-cover:
	@echo "Testing with coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

## Clean the binary
clean:
	@echo "Cleaning..."
	@ -f main

uplib:
	@echo "Upgrading libraries..."
	@go get -u github.com/jindasoft/jinda-platform@latest
	@go mod tidy

uplibdev:
	@echo "Upgrading libraries..."
	@go get -u github.com/jindasoft/jinda-platform@develop
	@go mod tidy

vuln:
	@echo "Checking for vulnerabilities..."
	@govulncheck ./...

vuln-v:
	@echo "Checking for vulnerabilities (verbose)..."
	@govulncheck -show=verbose ./...

sec:
	@echo "Checking for security issues..."
	@gosec ./...

swag:
	@echo "Generating Swagger docs..."
	@swag init --parseDependency --parseInternal -g cmd/api/main.go

trivy:
	@echo "Running Trivy scan..."
	@trivy fs \
		--db-repository ghcr.io/aquasecurity/trivy-db \
		--timeout 15m \
		--exit-code 1 \
		--severity HIGH,CRITICAL \
		--skip-files configs/secret.json .

check:
	@echo "Pre-commit check..."
	@pre-commit run --all-files
