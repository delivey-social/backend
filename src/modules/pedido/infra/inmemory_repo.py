from uuid import UUID

from src.modules.pedido.domain.pedido import Pedido
from src.modules.pedido.domain.repository import PedidoRepository


class InMemoryPedidoRepository(PedidoRepository):
    def __init__(self):
        self.pedidos: dict[UUID, Pedido] = {}

    async def list_pedidos(self) -> list[Pedido]:
        return list(self.pedidos.values())

    async def save(self, pedido: Pedido):
        self.pedidos[pedido.id] = pedido

    async def get_pedido(self, pedido_id: UUID) -> Pedido | None:
        return self.pedidos.get(pedido_id)
