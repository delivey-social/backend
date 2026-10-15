from uuid import UUID


class RestauranteService:
    def __init__(self):
        pass

    async def list_restaurantes(self) -> list[str]:
        print("Listing restaurantes")

        return []

    async def get_restaurante(self, restaurante_id: UUID):
        print(f"Getting restaurante with ID: {restaurante_id}")

        return {"id": str(restaurante_id), "name": "Sample Restaurante"}

    async def create_restaurante(self):
        print("Creating a new restaurante")

    async def update_restaurante(self, restaurante_id: UUID):
        print(f"Updating restaurante with ID: {restaurante_id}")

    async def delete_restaurante(self, restaurante_id: UUID):
        print(f"Deleting restaurante with ID: {restaurante_id}")
