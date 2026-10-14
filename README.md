## Development

Install dependencies with `uv sync`, then start the API with:

```bash
uv run uvicorn src.main:app --reload
```

The API is available at `http://127.0.0.1:8000`. Interactive OpenAPI
documentation is available at `/docs`.