from uuid import UUID

from fastapi import FastAPI

from src.modules.pedido.types import RestaurantGateway, UserGateway

from .modules.pedido.presentation import PedidoRouter
from .modules.pedido.service import PedidoService

app = FastAPI(title="Comida App Backend")


class DummyUserGateway(UserGateway):
    def exists(self, user_id: UUID) -> bool:
        return True


class DummyRestaurantGateway(RestaurantGateway):
    def exists(self, restaurant_id: UUID) -> bool:
        return True

    def item_exists(self, restaurant_id: UUID, product_id: UUID) -> bool:
        return True


pedido_service = PedidoService(DummyUserGateway(), DummyRestaurantGateway())
pedido_router = PedidoRouter(pedido_service)

app.include_router(pedido_router)


@app.get("/")
async def ping() -> dict[str, str]:
    return {"message": "comida.app backend is alive!"}
