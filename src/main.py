from collections.abc import AsyncGenerator, AsyncIterator
from contextlib import asynccontextmanager
from typing import Annotated
from uuid import UUID

from fastapi import Depends, FastAPI, Request
from fastapi.responses import JSONResponse
from sqlalchemy.ext.asyncio import AsyncSession

from src.core.errors import DomainError
from src.core.infra.database import engine, get_session
from src.core.infra.sqlalchemy_base import Base
from src.modules.pedido.types import UserGateway
from src.modules.restaurante.infra.inmemory_menu_repo import InMemoryMenuRepository
from src.modules.restaurante.infra.sqlalchemy_repo import (
    SQLAlchemyRestauranteRepository,
)
from src.modules.restaurante.menu_service import MenuService
from src.modules.restaurante.presentation import MenuRouter, RestauranteRouter
from src.modules.restaurante.service import RestauranteService

from .modules.pedido.infra.gateway_impl import RestauranteGatewayImplementation
from .modules.pedido.presentation import PedidoRouter
from .modules.pedido.service import PedidoService


@asynccontextmanager
async def lifespan(_: FastAPI) -> AsyncGenerator[None, None]:
    async with engine.begin() as connection:
        await connection.run_sync(Base.metadata.create_all)
    yield
    await engine.dispose()


app = FastAPI(title="Comida App Backend", lifespan=lifespan)


class DummyUserGateway(UserGateway):
    async def exists(self, user_id: UUID) -> bool:
        return True


SessionDependency = Annotated[AsyncSession, Depends(get_session)]


menu_repo = InMemoryMenuRepository()


async def get_restaurante_service(
    session: SessionDependency,
) -> AsyncIterator[RestauranteService]:
    try:
        yield RestauranteService(SQLAlchemyRestauranteRepository(session), menu_repo)
        await session.commit()
    except Exception:
        await session.rollback()
        raise


async def get_pedido_service(
    restaurante_service: Annotated[
        RestauranteService, Depends(get_restaurante_service)
    ],
) -> PedidoService:
    return PedidoService(
        DummyUserGateway(),
        RestauranteGatewayImplementation(restaurante_service),
    )


menu_service = MenuService(menu_repo)
menu_router = MenuRouter(menu_service)

restaurant_router = RestauranteRouter(get_restaurante_service)
pedido_router = PedidoRouter(get_pedido_service)

app.include_router(pedido_router)
app.include_router(menu_router)
app.include_router(restaurant_router)


@app.get("/")
async def ping() -> dict[str, str]:
    return {"message": "comida.app backend is alive!"}


@app.exception_handler(DomainError)
async def domain_error_handler(request: Request, exc: DomainError):
    return JSONResponse(
        status_code=400,
        content={"message": str(exc)},
    )
