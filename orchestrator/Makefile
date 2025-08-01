include .env

DOMAIN=orchestrator.local
CERT_DIR=certs
CERT_FILE=$(CERT_DIR)/$(DOMAIN).pem
KEY_FILE=$(CERT_DIR)/$(DOMAIN)-key.pem

.PHONY: create-certs clean-certs

# Create certificates
create-certs:
	@if ! command -v mkcert >/dev/null 2>&1; then \
		echo "Mkcert not found. Please install it."; \
		exit 1; \
	fi
	@mkcert -install
	@mkdir -p $(CERT_DIR)
	@if [ ! -f "$(CERT_FILE)" ] || [ ! -f "$(KEY_FILE)" ]; then \
		echo "Generating TLS cert for $(DOMAIN)..."; \
		mkcert -cert-file $(CERT_FILE) -key-file $(KEY_FILE) $(DOMAIN); \
	else \
		echo "Certs already exist for $(DOMAIN)"; \
	fi

clean-certs:
	@echo "Removing certificates..."
	@rm -f $(CERT_FILE) $(KEY_FILE)
	@rmdir --ignore-fail-on-non-empty $(CERT_DIR) 2>/dev/null || true
	@echo "Certificates removed."

watch:
	air

gql-init:
	go run github.com/99designs/gqlgen init

gql-generate:
	go run github.com/99designs/gqlgen generate
