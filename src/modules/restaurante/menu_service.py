from uuid import UUID

from src.core.errors import DomainError

from .domain.repository import MenuRepository


class MenuService:
    def __init__(self, menu_repo: MenuRepository):
        self.menu_repo = menu_repo

    async def get_menu(self, restaurante_id: UUID):
        menu = await self.menu_repo.get_by_restaurante_id(restaurante_id)

        if not menu:
            raise DomainError("Menu not found for the given restaurante_id")

        return menu
