from uuid import UUID

from .domain.restaurante import Restaurante


class RestauranteService:
    def __init__(self):
        pass

    async def list_restaurantes(self) -> list[Restaurante]:
        print("Listing restaurantes")

        return [
            Restaurante(
                id=UUID("12345678-1234-1234-1234-123456789012"),
                nome="Sample Restaurante",
            )
        ]

    async def get_restaurante(self, restaurante_id: UUID) -> Restaurante:
        print(f"Getting restaurante with ID: {restaurante_id}")

        return Restaurante(
            id=restaurante_id,
            nome="Sample Restaurante",
        )

    async def create_restaurante(self) -> Restaurante:
        print("Creating a new restaurante")

        return Restaurante(
            id=UUID("12345678-1234-1234-1234-123456789012"),
            nome="Sample Restaurante",
        )

    async def update_restaurante(self, restaurante_id: UUID) -> Restaurante:
        print(f"Updating restaurante with ID: {restaurante_id}")

        return Restaurante(
            id=restaurante_id,
            nome="Sample Restaurante",
        )

    async def delete_restaurante(self, restaurante_id: UUID):
        print(f"Deleting restaurante with ID: {restaurante_id}")
