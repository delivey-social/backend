from uuid import UUID

from ..domain.repository import RestauranteRepository
from ..domain.restaurante import Restaurante


class InMemoryRestauranteRepository(RestauranteRepository):
    def __init__(self):
        self.restaurantes: dict[UUID, Restaurante] = {}

    async def list_(self) -> list[Restaurante]:
        return list(self.restaurantes.values())

    async def get_restaurante(self, restaurante_id: UUID):
        return self.restaurantes.get(restaurante_id)

    async def save(self, data: Restaurante):
        self.restaurantes[data.id] = data

    async def update(self, data: Restaurante):
        if data.id in self.restaurantes:
            self.restaurantes[data.id] = data

    async def delete(self, restaurante_id: UUID):
        if restaurante_id in self.restaurantes:
            del self.restaurantes[restaurante_id]

    async def get_by_id(self, restaurante_id: UUID) -> Restaurante | None:
        result = self.restaurantes.get(restaurante_id)

        return result
