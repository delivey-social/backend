from fastapi import FastAPI

app = FastAPI(title="Comida App Backend")


@app.get("/")
async def ping() -> dict[str, str]:
    return {"message": "comida.app backend is alive!"}
