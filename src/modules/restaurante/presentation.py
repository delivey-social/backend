from uuid import UUID

from fastapi import APIRouter


class RestauranteRouter(APIRouter):
    def __init__(self):
        super().__init__(prefix="/restaurante", tags=["Restaurante"])

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
        # Logic to list restaurantes goes here
        return {"message": "List of restaurantes"}

    async def get_restaurante(self, restaurante_id: UUID):
        # Logic to get a specific restaurante by ID goes here
        return {"message": f"Details of restaurante {restaurante_id}"}

    async def create_restaurante(self):
        # Logic to create a new restaurante goes here
        return {"message": "Restaurante criado"}

    async def update_restaurante(self, restaurante_id: UUID):
        # Logic to update a specific restaurante by ID goes here
        return {"message": f"Restaurante {restaurante_id} updated"}

    async def delete_restaurante(self, restaurante_id: UUID):
        # Logic to delete a specific restaurante by ID goes here
        return {"message": f"Restaurante {restaurante_id} deleted"}
