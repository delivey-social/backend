from src.modules.pedido.domain.pedido import Pedido, PedidoItem
from src.modules.pedido.types import CreatePedidoRequest, RestaurantGateway, UserGateway


class PedidoService:
    def __init__(
        self,
        user_gateway: UserGateway,
        restaurant_gateway: RestaurantGateway,
    ):
        self.user_gateway = user_gateway
        self.restaurant_gateway = restaurant_gateway

    async def create_pedido(self, data: CreatePedidoRequest):
        exists = await self.user_gateway.exists(data.user_id)

        if not exists:
            raise ValueError(f"User with ID {data.user_id} does not exist.")

        exists = await self.restaurant_gateway.exists(data.restaurant_id)

        if not exists:
            raise ValueError(f"Restaurant with ID {data.restaurant_id} does not exist.")

        exists = True

        for item in data.items:
            if not await self.restaurant_gateway.item_exists(
                data.restaurant_id, item.product_id
            ):
                exists = False
                break

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

        print(f"Creating pedido with data: {pedido}")
