from .domain.pedido import Pedido, PedidoItem
from .domain.repository import PedidoRepository
from .types import (
    CreatePedidoRequest,
    RestaurantGateway,
    UserGateway,
)


class PedidoService:
    def __init__(
        self,
        repo: PedidoRepository,
        user_gateway: UserGateway,
        restaurant_gateway: RestaurantGateway,
    ):
        self.repo = repo
        self.user_gateway = user_gateway
        self.restaurant_gateway = restaurant_gateway

    async def list_pedidos(self) -> list[Pedido]:
        pedidos = await self.repo.list_pedidos()

        return pedidos

    async def create_pedido(self, data: CreatePedidoRequest):
        print("[PedidoService] Creating pedido")

        exists = await self.user_gateway.exists(data.user_id)

        if not exists:
            raise ValueError(f"User with ID {data.user_id} does not exist.")

        exists = await self.restaurant_gateway.exists(data.restaurant_id)

        if not exists:
            raise ValueError(f"Restaurant with ID {data.restaurant_id} does not exist.")

        items: list[PedidoItem] = []

        for item in data.items:
            item_domain = await self.restaurant_gateway.get_item(
                data.restaurant_id, item.product_id
            )

            items.append(
                PedidoItem(
                    product_id=item_domain.product_id,
                    quantity=item.quantity,
                    price_snapshot=item_domain.price,
                )
            )

        if not exists:
            raise ValueError(
                f"One or more items do not exist in the restaurant {data.restaurant_id}."
            )

        pedido = Pedido(
            user_id=data.user_id,
            restaurant_id=data.restaurant_id,
            items=[
                PedidoItem(
                    product_id=item.product_id,
                    quantity=item.quantity,
                    price_snapshot=0.0,
                )
                for item in data.items
            ],
        )

        await self.repo.save(pedido)
