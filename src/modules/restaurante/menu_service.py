from pathlib import Path
from uuid import UUID, uuid4

from fastapi import UploadFile

from src.core.errors import DomainError
from src.modules.restaurante.domain.menu import MenuItem
from src.modules.restaurante.types import CreateMenuItemRequest

from .domain.menu_repository import MenuRepository
from .file_service import FileService


class MenuService:
    def __init__(
        self,
        menu_repo: MenuRepository,
        file_service: FileService,
    ):
        self.menu_repo = menu_repo
        self.file_service = file_service

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
            raise DomainError("Item não encontrado", 404)

        item.nome = data.nome
        item.preco = data.preco

        await self.menu_repo.update_item(restaurante_id, item)

        return item

    async def delete_menu_item(self, restaurante_id: UUID, item_id: UUID):
        menu = await self.menu_repo.get_by_restaurante_id(restaurante_id)

        if not menu:
            raise DomainError("Menu não encontrado para o restaurante", 404)

        item = next((item for item in menu.itens if item.id == item_id), None)

        if not item:
            raise DomainError("Item não encontrado", 404)

        await self.menu_repo.delete_item(restaurante_id, item_id)

    async def get_menu_item(self, restaurante_id: UUID, item_id: UUID) -> MenuItem:
        item = await self.menu_repo.get_item_by_id(restaurante_id, item_id)

        if not item:
            raise DomainError("Item não encontrado", 404)

        return item

    async def upload_menu_item_image(
        self,
        restaurante_id: UUID,
        item_id: UUID,
        file: UploadFile,
    ):
        print(f"[MenuService] Uploading image for menu item with ID: {item_id}")

        exists = bool(
            await self.menu_repo.get_item_by_id(
                restaurante_id,
                item_id,
            )
        )

        if not exists:
            raise DomainError("Item não encontrado", 404)

        await self.file_service.upload_image(
            file=file,
            base_path=Path("menu_items"),
            filename_override=str(item_id),
        )

    async def get_menu_item_image(
        self,
        restaurante_id: UUID,
        item_id: UUID,
    ) -> Path:
        return await self.file_service.get_image(
            base_path=Path("menu_items"),
            filename=str(item_id),
        )
