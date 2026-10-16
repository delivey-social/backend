from typing import Any
from uuid import UUID

from fastapi import FastAPI
from fastapi.responses import JSONResponse

from src.core.errors import DomainError
from src.modules.pedido.types import UserGateway
from src.modules.restaurante.infra.inmemory_menu_repo import InMemoryMenuRepository
from src.modules.restaurante.infra.inmemory_repo import InMemoryRestauranteRepository
from src.modules.restaurante.menu_service import MenuService
from src.modules.restaurante.presentation import MenuRouter, RestauranteRouter
from src.modules.restaurante.service import RestauranteService

from .modules.pedido.infra.gateway_impl import RestauranteGatewayImplementation
from .modules.pedido.presentation import PedidoRouter
from .modules.pedido.service import PedidoService

app = FastAPI(title="Comida App Backend")


class DummyUserGateway(UserGateway):
    async def exists(self, user_id: UUID) -> bool:
        return True


menu_repo = InMemoryMenuRepository()
menu_service = MenuService(menu_repo)
menu_router = MenuRouter(menu_service)

restaurante_repo = InMemoryRestauranteRepository()
restaurant_service = RestauranteService(restaurante_repo, menu_repo)
restaurant_router = RestauranteRouter(restaurant_service)


pedido_service = PedidoService(
    DummyUserGateway(),
    RestauranteGatewayImplementation(restaurant_service),
)
pedido_router = PedidoRouter(pedido_service)

app.include_router(pedido_router)
app.include_router(menu_router)
app.include_router(restaurant_router)


@app.get("/")
async def ping() -> dict[str, str]:
    return {"message": "comida.app backend is alive!"}


@app.exception_handler(DomainError)
async def domain_error_handler(request: Any, exc: DomainError):
    return JSONResponse(
        status_code=400,
        content={"message": str(exc)},
    )
