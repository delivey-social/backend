.PHONY: dev tests pre-commit build start typecheck

dev:
	uv run fastapi dev src/main.py

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
	uv run ruff check
	uv run pyright
