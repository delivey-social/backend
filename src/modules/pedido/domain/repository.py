from abc import ABC, abstractmethod

from .pedido import Pedido


class PedidoRepository(ABC):
    @abstractmethod
    async def list_pedidos(self) -> list[Pedido]:
        pass

    @abstractmethod
    async def save(self, pedido: Pedido):
        pass
