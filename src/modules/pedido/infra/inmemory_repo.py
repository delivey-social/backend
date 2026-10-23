from src.modules.pedido.domain.pedido import Pedido
from src.modules.pedido.domain.repository import PedidoRepository


class InMemoryPedidoRepository(PedidoRepository):
    def __init__(self):
        self.pedidos: list[Pedido] = []

    async def list_pedidos(self) -> list[Pedido]:
        return self.pedidos
