from uuid import UUID

from sqlalchemy import select
from sqlalchemy.ext.asyncio import AsyncSession

from src.modules.restaurante.domain.menu import Menu, MenuItem
from src.modules.restaurante.domain.menu_repository import MenuRepository
from src.modules.restaurante.infra.sqlalchemy_model import MenuItemModel, MenuModel


class SQLAlchemyMenuRepository(MenuRepository):
    def __init__(self, session: AsyncSession):
        self.session = session

    async def save(self, restaurante_id: UUID):
        menu = MenuModel(restaurante_id=restaurante_id)

        self.session.add(menu)

    async def get_by_restaurante_id(
        self,
        restaurante_id: UUID,
    ) -> Menu | None:
        result = await self.session.execute(
            select(MenuModel).where(MenuModel.restaurante_id == restaurante_id)
        )
        model = result.scalar_one_or_none()

        if model is None:
            return None

        return Menu(
            id=model.id,
            restaurante_id=model.restaurante_id,
            itens=[
                MenuItem(id=item.id, nome=item.nome, preco=item.preco)
                for item in model.itens
            ],
        )

    async def create_item(self, restaurante_id: UUID, data: MenuItem):
        result = await self.session.execute(
            select(MenuModel).where(MenuModel.restaurante_id == restaurante_id)
        )
        menu_model = result.scalar_one_or_none()

        if menu_model is None:
            raise ValueError("Menu not found for the given restaurante_id")

        menu_item_model = MenuItemModel(
            id=data.id, nome=data.nome, preco=data.preco, menu=menu_model
        )

        self.session.add(menu_item_model)

    async def get_item_by_id(
        self, restaurante_id: UUID, item_id: UUID
    ) -> MenuItem | None:
        stmt = (
            select(MenuItemModel)
            .join(
                MenuItemModel.menu
            )  # Join through the relationship defined in the model
            .where(
                MenuItemModel.id == item_id, MenuModel.restaurante_id == restaurante_id
            )
        )

        result = await self.session.execute(stmt)
        item_model = result.scalar_one_or_none()

        if item_model is None:
            return None

        return MenuItem(
            id=item_model.id,
            nome=item_model.nome,
            preco=item_model.preco,
        )

    async def update_item(self, restaurante_id: UUID, data: MenuItem):
        stmt = (
            select(MenuItemModel)
            .join(MenuItemModel.menu)
            .where(
                MenuItemModel.id == data.id, MenuModel.restaurante_id == restaurante_id
            )
        )

        result = await self.session.execute(stmt)
        item_model = result.scalar_one_or_none()

        if item_model is None:
            raise ValueError("Menu item not found for the given item_id")

        item_model.nome = data.nome
        item_model.preco = data.preco

    async def delete_item(self, restaurante_id: UUID, item_id: UUID):
        stmt = (
            select(MenuItemModel)
            .join(MenuItemModel.menu)
            .where(
                MenuItemModel.id == item_id, MenuModel.restaurante_id == restaurante_id
            )
        )

        result = await self.session.execute(stmt)
        item_model = result.scalar_one_or_none()

        if item_model is None:
            raise ValueError("Menu item not found for the given item_id")

        await self.session.delete(item_model)
