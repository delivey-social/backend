from uuid import UUID, uuid4

from src.core.errors import DomainError
from src.modules.restaurante.types import (
    CreateRestauranteRequest,
    UpdateRestauranteRequest,
)

from .domain.repository import MenuRepository, RestauranteRepository
from .domain.restaurante import Restaurante


class RestauranteService:
    def __init__(self, repo: RestauranteRepository, menu_repo: MenuRepository):
        self.repo = repo
        self.menu_repo = menu_repo

    async def list_restaurantes(self) -> list[Restaurante]:
        print("[RestauranteService] Listando restaurantes...")

        return await self.repo.list_restaurantes()

    async def get_restaurante(self, restaurante_id: UUID) -> Restaurante:
        print(f"[RestauranteService] Obtendo restaurante com ID: {restaurante_id}")

        restaurant = await self.repo.get_restaurante_by_id(restaurante_id)

        if not restaurant:
            raise DomainError(
                f"Restaurante with ID {restaurante_id} not found",
            )

        return restaurant

    async def create_restaurante(
        self,
        data: CreateRestauranteRequest,
    ) -> Restaurante:
        print("[RestauranteService] Criando novo restaurante...")

        restaurante = Restaurante(
            id=uuid4(),
            nome=data.nome,
        )

        await self.repo.create_restaurante(restaurante)
        await self.menu_repo.create_menu(restaurante.id)

        return restaurante

    async def update_restaurante(
        self,
        data: UpdateRestauranteRequest,
    ) -> Restaurante:
        print(f"[RestauranteService] Atualizando restaurante com ID: {data.id}")

        restaurante = Restaurante(
            id=data.id,
            nome=data.nome,
        )

        await self.repo.update_restaurante(restaurante)

        return restaurante

    async def delete_restaurante(self, restaurante_id: UUID):
        print(f"[RestauranteService] Excluindo restaurante com ID: {restaurante_id}")

        return await self.repo.delete_restaurante(restaurante_id)
