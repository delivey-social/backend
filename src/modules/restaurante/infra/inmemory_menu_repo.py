from uuid import UUID, uuid4

from ..domain.menu import Menu, MenuItem
from ..domain.menu_repository import MenuRepository


class InMemoryMenuRepository(MenuRepository):
    def __init__(self):
        self.menus: dict[UUID, Menu] = {}

    async def save(self, restaurante_id: UUID):
        self.menus[restaurante_id] = Menu(
            id=uuid4(),
            restaurante_id=restaurante_id,
            itens=[],
        )

    async def get_by_restaurante_id(self, restaurante_id: UUID) -> Menu | None:
        return self.menus.get(restaurante_id)

    async def create_item(self, restaurante_id: UUID, data: MenuItem):
        menu = self.menus.get(restaurante_id)

        if menu is None:
            raise ValueError("Menu not found for the given restaurante_id")

        menu.itens.append(data)
