from uuid import UUID, uuid4

from src.core.errors import DomainError
from src.modules.restaurante.domain.menu import MenuItem
from src.modules.restaurante.types import CreateMenuItemRequest

from .domain.menu_repository import MenuRepository


class MenuService:
    def __init__(self, menu_repo: MenuRepository):
        self.menu_repo = menu_repo

    async def get_menu(self, restaurante_id: UUID):
        menu = await self.menu_repo.get_by_restaurante_id(restaurante_id)

        if not menu:
            raise DomainError("Menu não encontrado para o restaurante", 404)

        return menu

    async def create_menu_item(self, restaurante_id: UUID, data: CreateMenuItemRequest):
        await self.menu_repo.create_item(
            restaurante_id,
            MenuItem(
                id=uuid4(),
                nome=data.nome,
                preco=data.preco,
            ),
        )
