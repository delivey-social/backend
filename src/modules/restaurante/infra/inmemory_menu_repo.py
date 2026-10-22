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

    async def get_item_by_id(
        self, restaurante_id: UUID, item_id: UUID
    ) -> MenuItem | None:
        menu = self.menus.get(restaurante_id)

        if menu is None:
            return None

        return next((item for item in menu.itens if item.id == item_id), None)

    async def update_item(self, restaurante_id: UUID, data: MenuItem):
        menu = self.menus.get(restaurante_id)

        if menu is None:
            raise ValueError("Menu not found for the given restaurante_id")

        item = await self.get_item_by_id(restaurante_id, data.id)

        if item is None:
            raise ValueError("Menu item not found for the given item_id")

        item.nome = data.nome
        item.preco = data.preco

    async def delete_item(self, restaurante_id: UUID, item_id: UUID):
        menu = self.menus.get(restaurante_id)

        if menu is None:
            raise ValueError("Menu not found for the given restaurante_id")

        item = await self.get_item_by_id(restaurante_id, item_id)

        if item is None:
            raise ValueError("Menu item not found for the given item_id")

        menu.itens.remove(item)
