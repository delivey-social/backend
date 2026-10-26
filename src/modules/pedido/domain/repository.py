from abc import ABC, abstractmethod
from uuid import UUID

from .pedido import Pedido


class PedidoRepository(ABC):
    @abstractmethod
    async def list_pedidos(self) -> list[Pedido]:
        pass

    @abstractmethod
    async def save(self, pedido: Pedido):
        pass

    @abstractmethod
    async def get_pedido(self, pedido_id: UUID) -> Pedido | None:
        pass

    @abstractmethod
    async def delete_pedido(self, pedido_id: UUID):
        pass

    @abstractmethod
    async def update_pedido(self, pedido: Pedido):
        pass
