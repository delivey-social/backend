from uuid import UUID

from fastapi import FastAPI

from src.modules.pedido.types import RestaurantGateway, UserGateway

from .modules.pedido.presentation import PedidoRouter
from .modules.pedido.service import PedidoService
from .modules.pedido.types import RestaurantItem

app = FastAPI(title="Comida App Backend")


class DummyUserGateway(UserGateway):
    async def exists(self, user_id: UUID) -> bool:
        return True


class DummyRestaurantGateway(RestaurantGateway):
    async def exists(self, restaurant_id: UUID) -> bool:
        return True

    async def get_item(self, restaurant_id: UUID, product_id: UUID) -> RestaurantItem:
        return RestaurantItem(
            product_id=product_id,
            name="Dummy Item",
            price=100,
        )


pedido_service = PedidoService(DummyUserGateway(), DummyRestaurantGateway())
pedido_router = PedidoRouter(pedido_service)

app.include_router(pedido_router)


@app.get("/")
async def ping() -> dict[str, str]:
    return {"message": "comida.app backend is alive!"}
