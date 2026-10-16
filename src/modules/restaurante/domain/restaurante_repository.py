from abc import ABC, abstractmethod
from uuid import UUID

from .restaurante import Restaurante


class RestauranteRepository(ABC):
    @abstractmethod
    async def list_(self) -> list[Restaurante]:
        pass

    @abstractmethod
    async def get_by_id(self, restaurante_id: UUID) -> Restaurante | None:
        pass

    @abstractmethod
    async def save(self, data: Restaurante) -> None:
        pass

    @abstractmethod
    async def update(self, data: Restaurante) -> None:
        pass

    @abstractmethod
    async def delete(self, restaurante_id: UUID) -> None:
        pass
