from collections.abc import AsyncIterator
from typing import Annotated

from fastapi import Depends
from sqlalchemy.ext.asyncio import AsyncSession

from src.core.infra.database import get_session
from src.modules.restaurante.infra.sqlalchemy_repo import (
    SQLAlchemyRestauranteRepository,
)
from src.modules.restaurante.menu_service import MenuService
from src.modules.restaurante.service import RestauranteService

from .inmemory_menu_repo import InMemoryMenuRepository

SessionDependency = Annotated[AsyncSession, Depends(get_session)]

menu_repo = InMemoryMenuRepository()
menu_service = MenuService(menu_repo)


async def get_restaurante_service(
    session: SessionDependency,
) -> AsyncIterator[RestauranteService]:
    try:
        yield RestauranteService(SQLAlchemyRestauranteRepository(session), menu_repo)
        await session.commit()
    except Exception:
        await session.rollback()
        raise
