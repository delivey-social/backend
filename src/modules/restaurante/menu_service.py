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

    async def create_menu_item(
        self,
        restaurante_id: UUID,
        data: CreateMenuItemRequest,
    ) -> MenuItem:
        menu = await self.menu_repo.get_by_restaurante_id(restaurante_id)

        if not menu:
            raise DomainError("Menu não encontrado para o restaurante", 404)

        item = MenuItem(
            id=uuid4(),
            nome=data.nome,
            preco=data.preco,
        )

        await self.menu_repo.create_item(restaurante_id, item)

        return item

    async def update_menu_item(
        self,
        restaurante_id: UUID,
        item_id: UUID,
        data: CreateMenuItemRequest,
    ) -> MenuItem:
        menu = await self.menu_repo.get_by_restaurante_id(restaurante_id)

        if not menu:
            raise DomainError("Menu não encontrado para o restaurante", 404)

        item = next((item for item in menu.itens if item.id == item_id), None)

        if not item:
            raise DomainError("Item do menu não encontrado", 404)

        item.nome = data.nome
        item.preco = data.preco

        await self.menu_repo.update_item(restaurante_id, item)

        return item
