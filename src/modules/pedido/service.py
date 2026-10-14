from src.modules.pedido.types import CreatePedidoRequest, RestaurantGateway, UserGateway


class PedidoService:
    def __init__(
        self,
        user_gateway: UserGateway,
        restaurant_gateway: RestaurantGateway,
    ):
        self.user_gateway = user_gateway
        self.restaurant_gateway = restaurant_gateway

    def create_pedido(self, data: CreatePedidoRequest):
        exists = self.user_gateway.exists(data.user_id)

        if not exists:
            raise ValueError(f"User with ID {data.user_id} does not exist.")

        exists = self.restaurant_gateway.exists(data.restaurant_id)

        if not exists:
            raise ValueError(f"Restaurant with ID {data.restaurant_id} does not exist.")

        exists = all(
            self.restaurant_gateway.item_exists(data.restaurant_id, item.product_id)
            for item in data.items
        )

        if not exists:
            raise ValueError(
                f"One or more items do not exist in the restaurant {data.restaurant_id}."
            )

        print(f"Creating pedido with data: {data}")
