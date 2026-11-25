.PHONY: dev-backend keycloak-prod tests pre-commit install-hooks build start typecheck

dev-backend:
	@echo "This script is intended to be run from VSCode tasks."
	uv run python -m fastapi dev src/main.py

keycloak-prod:
	ENV=production docker compose up -d keycloak-db keycloak

tests:
	uv run python -m pytest tests

pre-commit:
	uv run python -m pre_commit run --all-files

install-hooks:
	uv run python -m pre_commit install

build:
	uv build
	uv pip install -e .

start:
	.venv/bin/start

typecheck:
	uv run pyright
