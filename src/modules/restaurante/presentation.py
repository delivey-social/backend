from uuid import UUID

from fastapi import APIRouter

from src.modules.restaurante.types import (
    CreateRestauranteRequest,
    UpdateRestauranteRequest,
)

from .service import RestauranteService


class RestauranteRouter(APIRouter):
    def __init__(self, service: RestauranteService):
        super().__init__(prefix="/restaurante", tags=["Restaurante"])

        self.service = service

        self.add_api_route(
            "/",
            self.list_restaurantes,
            methods=["GET"],
        )
        self.add_api_route(
            "/{restaurante_id}",
            self.get_restaurante,
            methods=["GET"],
        )

        self.add_api_route(
            "/",
            self.create_restaurante,
            methods=["POST"],
        )

        self.add_api_route(
            "/{restaurante_id}",
            self.update_restaurante,
            methods=["PUT"],
        )

        self.add_api_route(
            "/{restaurante_id}",
            self.delete_restaurante,
            methods=["DELETE"],
        )

    async def list_restaurantes(self):
        return await self.service.list_restaurantes()

    async def get_restaurante(self, restaurante_id: UUID):
        return await self.service.get_restaurante(restaurante_id)

    async def create_restaurante(self, data: CreateRestauranteRequest):
        await self.service.create_restaurante(data)

        return {"message": "Restaurante created successfully"}

    async def update_restaurante(
        self, restaurante_id: UUID, data: CreateRestauranteRequest
    ):
        await self.service.update_restaurante(
            UpdateRestauranteRequest(
                id=restaurante_id,
                nome=data.nome,
            )
        )

        return {"message": f"Restaurante {restaurante_id} updated"}

    async def delete_restaurante(self, restaurante_id: UUID):
        await self.service.delete_restaurante(restaurante_id)

        return {"message": f"Restaurante {restaurante_id} deleted"}
