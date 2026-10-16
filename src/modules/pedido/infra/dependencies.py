from typing import Annotated
from uuid import UUID

from fastapi import Depends

from src.modules.pedido.service import PedidoService
from src.modules.pedido.types import UserGateway
from src.modules.restaurante.infra.dependencies import get_restaurante_service
from src.modules.restaurante.service import RestauranteService

from .gateway_impl import RestauranteGatewayImplementation


class DummyUserGateway(UserGateway):
    async def exists(self, user_id: UUID) -> bool:
        return True


async def get_pedido_service(
    restaurante_service: Annotated[
        RestauranteService, Depends(get_restaurante_service)
    ],
) -> PedidoService:
    return PedidoService(
        DummyUserGateway(),
        RestauranteGatewayImplementation(restaurante_service),
    )
