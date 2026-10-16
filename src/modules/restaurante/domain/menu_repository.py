from abc import ABC, abstractmethod
from uuid import UUID

from .menu import Menu, MenuItem


class MenuRepository(ABC):
    @abstractmethod
    async def save(self, restaurante_id: UUID) -> None:
        pass

    @abstractmethod
    async def get_by_restaurante_id(self, restaurante_id: UUID) -> Menu | None:
        pass

    @abstractmethod
    async def create_item(self, restaurante_id: UUID, data: MenuItem) -> None:
        pass
