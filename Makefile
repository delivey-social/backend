.PHONY: dev-backend keycloak-prod tests pre-commit build start typecheck

dev-backend:
	@echo "This script is intended to be run from VSCode tasks."
	uv run fastapi dev src/main.py

keycloak-prod:
	ENV=production docker compose up -d keycloak-db keycloak

tests:
	uv run pytest tests

pre-commit:
	uv run pre-commit run --all-files

build:
	uv build
	uv pip install -e .

start:
	.venv/bin/start

typecheck:
	uv run pyright
