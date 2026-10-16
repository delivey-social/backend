from uuid import UUID

from src.modules.restaurante.service import RestauranteService

from ..types import RestaurantGateway, RestaurantItem


class RestauranteGatewayImplementation(RestaurantGateway):
    def __init__(self, service: RestauranteService):
        self.service = service

    async def exists(self, restaurant_id: UUID) -> bool:
        await self.service.get_restaurante(restaurant_id)

        return True

    async def get_item(self, restaurant_id: UUID, product_id: UUID):
        return RestaurantItem(
            product_id=product_id,
            name="Product Name",
            price=1000,
        )
