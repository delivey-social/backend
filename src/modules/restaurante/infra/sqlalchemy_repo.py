from uuid import UUID

from sqlalchemy import delete, select
from sqlalchemy.ext.asyncio import AsyncSession

from src.modules.restaurante.infra.sqlalchemy_model import RestauranteModel

from ..domain.repository import RestauranteRepository
from ..domain.restaurante import Restaurante


class SQLAlchemyRestauranteRepository(RestauranteRepository):
    def __init__(self, session: AsyncSession):
        self.session = session

    async def list_restaurantes(self) -> list[Restaurante]:
        result = await self.session.execute(select(RestauranteModel))

        return [
            Restaurante(id=model.id, nome=model.nome)
            for model in result.scalars().all()
        ]

    async def create_restaurante(self, data: Restaurante) -> None:
        self.session.add(
            RestauranteModel(
                id=data.id,
                nome=data.nome,
            )
        )

    async def get_restaurante_by_id(
        self,
        restaurante_id: UUID,
    ) -> Restaurante | None:
        model = await self.session.get(RestauranteModel, restaurante_id)

        if model is None:
            return None

        return Restaurante(id=model.id, nome=model.nome)

    async def update_restaurante(self, data: Restaurante) -> None:
        model = await self.session.get(RestauranteModel, data.id)

        if model is not None:
            model.nome = data.nome

    async def delete_restaurante(self, restaurante_id: UUID) -> None:
        await self.session.execute(
            delete(RestauranteModel).where(RestauranteModel.id == restaurante_id)
        )
