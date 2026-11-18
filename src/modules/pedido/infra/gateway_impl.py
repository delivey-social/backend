from uuid import UUID

from src.modules.restaurante.menu_service import MenuService
from src.modules.restaurante.service import RestauranteService

from ..types import RestaurantGateway, RestaurantItem


class RestauranteGatewayImplementation(RestaurantGateway):
    def __init__(
        self,
        service: RestauranteService,
        menu_service: MenuService,
    ):
        self.service = service
        self.menu_service = menu_service

    async def exists(self, restaurant_id: UUID) -> bool:
        await self.service.get_restaurante(restaurant_id)

        return True

    async def get_item(self, restaurant_id: UUID, product_id: UUID):
        item = await self.menu_service.get_menu_item(
            restaurant_id,
            product_id,
        )

        return RestaurantItem(
            product_id=item.id,
            name=item.nome,
            price=item.preco,
        )
