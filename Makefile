CODE_COVERAGE_DIR = test-results
export GOCOVERDIR := $(CODE_COVERAGE_DIR)

BUILD_DIR = ./bin
BINARY_NAME = suricata-listener

# Download dependencies
dep:
	go mod tidy && go mod vendor
	go list -mod=readonly -deps -f '{{define "M"}}{{.Path}} {{.Version}}{{end}}{{with .Module}}{{if not .Main}}{{if .Replace}}{{template "M" .Replace}}{{else}}{{template "M" .}}{{end}}{{end}}{{end}}' all | sort -u > go.list

# Setup necessary tools
setup: install

# Install necessary tools
install:
	@which ginkgo > /dev/null || go install github.com/onsi/ginkgo/v2/ginkgo@v2.8.0
	@which golangci-lint > /dev/null || go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.2
	@which goimports > /dev/null || go install golang.org/x/tools/cmd/goimports@v0.8.0

# Clean built and test files
clean:
	@rm -rf $(BUILD_DIR) ${CODE_COVERAGE_DIR} junit.xml coverprofile.out

# Ginkgo unit tests
test/unit:
	ginkgo -mod=vendor --vv --cover --fail-on-pending --junit-report=junit.xml \
		$(TEST_ARGS) \
		./...

# Run unit tests
test: test/unit

# Generate code coverage
coverage:
	rm -f ${CODE_COVERAGE_DIR}/coverage.*
	@$(MAKE) TEST_ARGS="--covermode=set --output-dir=${CODE_COVERAGE_DIR} --coverprofile=coverage.out" test/unit
	mkdir -p coverage_report/
	go tool cover -html ${CODE_COVERAGE_DIR}/coverage.out -o coverage_report/index.html

# Run linter
lint:
	golangci-lint run

# Build the binary
build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME)

# Run the application
run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

# Help information
help:
	@echo "  dep          Download dependencies and update go.list"
	@echo "  setup        Setup necessary tools for development"
	@echo "  install      Install required tools for development"
	@echo "  clean        Remove built and test files"
	@echo "  test/unit    Run Ginkgo unit tests"
	@echo "  test         Run unit tests"
	@echo "  coverage     Generate code coverage"
	@echo "  lint         Run golangci-lint"
	@echo "  build        Compile the binary"
	@echo "  run          Run the application"
