from fastapi import FastAPI

from .modules.pedido.presentation import PedidoRouter

app = FastAPI(title="Comida App Backend")

pedido_router = PedidoRouter()

app.include_router(pedido_router)


@app.get("/")
async def ping() -> dict[str, str]:
    return {"message": "comida.app backend is alive!"}
