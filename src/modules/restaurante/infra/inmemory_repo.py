from uuid import UUID

from ..domain.repository import RestauranteRepository
from ..domain.restaurante import Restaurante


class InMemoryRestauranteRepository(RestauranteRepository):
    def __init__(self):
        self.restaurantes: dict[UUID, Restaurante] = {}

    async def list_restaurantes(self) -> list[Restaurante]:
        return list(self.restaurantes.values())

    async def get_restaurante(self, restaurante_id: UUID):
        return self.restaurantes.get(restaurante_id)

    async def create_restaurante(self, data: Restaurante):
        self.restaurantes[data.id] = data

    async def update_restaurante(self, data: Restaurante):
        if data.id in self.restaurantes:
            self.restaurantes[data.id] = data

    async def delete_restaurante(self, restaurante_id: UUID):
        if restaurante_id in self.restaurantes:
            del self.restaurantes[restaurante_id]

    async def get_restaurante_by_id(self, restaurante_id: UUID) -> Restaurante:
        result = self.restaurantes.get(restaurante_id)

        if not result:
            raise ValueError("Restaurante not found")

        return result
