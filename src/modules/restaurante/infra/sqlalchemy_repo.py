from uuid import UUID

from sqlalchemy import delete, select
from sqlalchemy.ext.asyncio import AsyncSession

from src.modules.restaurante.infra.sqlalchemy_model import RestauranteModel

from ..domain.restaurante import Restaurante
from ..domain.restaurante_repository import RestauranteRepository


class SQLAlchemyRestauranteRepository(RestauranteRepository):
    def __init__(self, session: AsyncSession):
        self.session = session

    async def list_(self) -> list[Restaurante]:
        result = await self.session.execute(select(RestauranteModel))

        return [
            Restaurante(id=model.id, nome=model.nome)
            for model in result.scalars().all()
        ]

    async def save(self, data: Restaurante) -> None:
        self.session.add(
            RestauranteModel(
                id=data.id,
                nome=data.nome,
            )
        )

    async def get_by_id(
        self,
        restaurante_id: UUID,
    ) -> Restaurante | None:
        model = await self.session.get(RestauranteModel, restaurante_id)

        if model is None:
            return None

        return Restaurante(id=model.id, nome=model.nome)

    async def update(self, data: Restaurante) -> None:
        model = await self.session.get(RestauranteModel, data.id)

        if model is not None:
            model.nome = data.nome

    async def delete(self, restaurante_id: UUID) -> None:
        await self.session.execute(
            delete(RestauranteModel).where(RestauranteModel.id == restaurante_id)
        )
