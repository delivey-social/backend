from collections.abc import AsyncIterator
from typing import Annotated

from fastapi import Depends
from sqlalchemy.ext.asyncio import AsyncSession

from src.core.infra.database import get_session
from src.modules.restaurante.infra.disk_file_repo import DiskFileRepository

from ..domain.image_repository import FileRepository
from ..file_service import FileService
from ..infra.sqlalchemy_menu_repo import SQLAlchemyMenuRepository
from ..infra.sqlalchemy_repo import (
    SQLAlchemyRestauranteRepository,
)
from ..menu_service import MenuService
from ..service import RestauranteService

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


async def get_file_repo(
    session: SessionDependency,
) -> AsyncIterator[FileRepository]:
    try:
        yield DiskFileRepository()
        await session.commit()
    except Exception:
        await session.rollback()
        raise


async def get_file_service(
    session: SessionDependency,
    file_repo: Annotated[FileRepository, Depends(get_file_repo)],
) -> AsyncIterator[FileService]:
    try:
        yield FileService(file_repo)
        await session.commit()
    except Exception:
        await session.rollback()
        raise


async def get_menu_service(
    session: SessionDependency,
    menu_repo: Annotated[SQLAlchemyMenuRepository, Depends(get_menu_repo)],
    file_service: Annotated[FileService, Depends(get_file_service)],
) -> AsyncIterator[MenuService]:
    try:
        yield MenuService(menu_repo, file_service)
        await session.commit()
    except Exception:
        await session.rollback()
        raise


async def get_restaurante_service(
    session: SessionDependency,
    menu_repo: Annotated[SQLAlchemyMenuRepository, Depends(get_menu_repo)],
    file_service: Annotated[FileService, Depends(get_file_service)],
) -> AsyncIterator[RestauranteService]:
    try:
        yield RestauranteService(
            SQLAlchemyRestauranteRepository(session),
            menu_repo,
            file_service,
        )
        await session.commit()
    except Exception:
        await session.rollback()
        raise
