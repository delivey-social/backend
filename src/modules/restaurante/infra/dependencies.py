from collections.abc import AsyncIterator
from typing import Annotated

from fastapi import Depends
from sqlalchemy.ext.asyncio import AsyncSession

from src.core.infra.database import get_session
from src.modules.restaurante.infra.sqlalchemy_menu_repo import SQLAlchemyMenuRepository
from src.modules.restaurante.infra.sqlalchemy_repo import (
    SQLAlchemyRestauranteRepository,
)
from src.modules.restaurante.menu_service import MenuService
from src.modules.restaurante.service import RestauranteService

SessionDependency = Annotated[AsyncSession, Depends(get_session)]


async def get_menu_repo(
    session: SessionDependency,
) -> AsyncIterator[SQLAlchemyMenuRepository]:
    try:
        yield SQLAlchemyMenuRepository(session)
        await session.commit()
    except Exception:
        await session.rollback()
        raise


async def get_menu_service(
    session: SessionDependency,
    menu_repo: Annotated[SQLAlchemyMenuRepository, Depends(get_menu_repo)],
) -> AsyncIterator[MenuService]:
    try:
        yield MenuService(menu_repo)
        await session.commit()
    except Exception:
        await session.rollback()
        raise


async def get_restaurante_service(
    session: SessionDependency,
    menu_repo: Annotated[SQLAlchemyMenuRepository, Depends(get_menu_repo)],
) -> AsyncIterator[RestauranteService]:
    try:
        yield RestauranteService(SQLAlchemyRestauranteRepository(session), menu_repo)
        await session.commit()
    except Exception:
        await session.rollback()
        raise
