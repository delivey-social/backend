from uuid import UUID, uuid4

from ..domain.repository import MenuRepository
from ..domain.restaurante import Menu


class InMemoryMenuRepository(MenuRepository):
    def __init__(self):
        self.menus: dict[UUID, Menu] = {}

    async def create_menu(self, restaurante_id: UUID):
        self.menus[restaurante_id] = Menu(
            id=uuid4(),
            restaurante_id=restaurante_id,
            itens=[],
        )

    async def get_menu_by_restaurante_id(self, restaurante_id: UUID) -> Menu | None:
        return self.menus.get(restaurante_id)
