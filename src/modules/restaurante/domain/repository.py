from abc import ABC, abstractmethod
from uuid import UUID

from .restaurante import Menu, Restaurante


class RestauranteRepository(ABC):
    @abstractmethod
    async def list_restaurantes(self) -> list[Restaurante]:
        pass

    @abstractmethod
    async def get_restaurante_by_id(self, restaurante_id: UUID) -> Restaurante | None:
        pass

    @abstractmethod
    async def create_restaurante(self, data: Restaurante):
        pass

    @abstractmethod
    async def update_restaurante(self, data: Restaurante):
        pass

    @abstractmethod
    async def delete_restaurante(self, restaurante_id: UUID):
        pass


class MenuRepository(ABC):
    @abstractmethod
    async def create_menu(self, restaurante_id: UUID):
        pass

    @abstractmethod
    async def get_menu_by_restaurante_id(self, restaurante_id: UUID) -> Menu | None:
        pass
