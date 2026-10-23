from collections.abc import AsyncIterator
from typing import Annotated
from uuid import UUID

from fastapi import Depends

from src.modules.pedido.domain.repository import PedidoRepository
from src.modules.restaurante.infra.dependencies import get_restaurante_service
from src.modules.restaurante.service import RestauranteService

from ..infra.inmemory_repo import InMemoryPedidoRepository
from ..service import PedidoService
from ..types import UserGateway
from .gateway_impl import RestauranteGatewayImplementation


pedido_repository = InMemoryPedidoRepository()


class DummyUserGateway(UserGateway):
    async def exists(self, user_id: UUID) -> bool:
        return True


async def get_pedido_repo() -> AsyncIterator[PedidoRepository]:
    yield pedido_repository


async def get_pedido_service(
    restaurante_service: Annotated[
        RestauranteService, Depends(get_restaurante_service)
    ],
    pedido_repository: Annotated[PedidoRepository, Depends(get_pedido_repo)],
) -> PedidoService:
    return PedidoService(
        pedido_repository,
        DummyUserGateway(),
        RestauranteGatewayImplementation(restaurante_service),
    )
