# Variables
BACKEND_DIR := backend
FRONTEND_DIR := frontend

# Targets
.PHONY: test-backend lint-backend check-backend build-frontend

test-backend:
	cd $(BACKEND_DIR) && \
	go test -v ./...
lint-backend:
	cd $(BACKEND_DIR) && \
	golangci-lint run

check-backend:
	cd $(BACKEND_DIR) && \
	go vet ./... && \

build-frontend:
	cd $(FRONTEND_DIR) && \
	npm run builds