from uuid import UUID

from src.modules.pedido.domain.pedido import Pedido
from src.modules.pedido.domain.repository import PedidoRepository


class InMemoryPedidoRepository(PedidoRepository):
    def __init__(self):
        self.pedidos: dict[UUID, Pedido] = {}

    async def list_pedidos(self) -> list[Pedido]:
        print("[InMemoryPedidoRepository] Listing all pedidos")

        return list(self.pedidos.values())

    async def save(self, pedido: Pedido):
        print(f"[InMemoryPedidoRepository] Saving pedido with ID {pedido.id}")

        self.pedidos[pedido.id] = pedido

    async def get_pedido(self, pedido_id: UUID) -> Pedido | None:
        print(f"[InMemoryPedidoRepository] Getting pedido with ID {pedido_id}")

        return self.pedidos.get(pedido_id)

    async def delete_pedido(self, pedido_id: UUID):
        print(f"[InMemoryPedidoRepository] Deleting pedido with ID {pedido_id}")

        if pedido_id in self.pedidos:
            del self.pedidos[pedido_id]
        else:
            print(f"[InMemoryPedidoRepository] Pedido with ID {pedido_id} not found")

            raise ValueError(f"Pedido with ID {pedido_id} does not exist.")

    async def update_pedido(self, pedido: Pedido):
        print(f"[InMemoryPedidoRepository] Updating pedido with ID {pedido.id}")

        if pedido.id in self.pedidos:
            self.pedidos[pedido.id] = pedido
        else:
            print(f"[InMemoryPedidoRepository] Pedido with ID {pedido.id} not found")

            raise ValueError(f"Pedido with ID {pedido.id} does not exist.")
