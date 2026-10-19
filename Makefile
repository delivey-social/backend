.PHONY: dev tests

dev:
	uv run fastapi dev src/main.py

tests:
	uv run pytest tests