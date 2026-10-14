from fastapi import FastAPI

from .modules.pedido.presentation import PedidoRouter
from .modules.pedido.service import PedidoService

app = FastAPI(title="Comida App Backend")

pedido_service = PedidoService()
pedido_router = PedidoRouter(pedido_service)

app.include_router(pedido_router)


@app.get("/")
async def ping() -> dict[str, str]:
    return {"message": "comida.app backend is alive!"}
