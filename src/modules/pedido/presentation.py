from uuid import UUID

from fastapi import APIRouter

from src.modules.pedido.service import PedidoService
from src.modules.pedido.types import CreatePedidoRequest


class PedidoRouter(APIRouter):
    def __init__(self, service: PedidoService):
        self.service = service

        super().__init__(prefix="/pedido", tags=["Pedido"])

        self.add_api_route(
            "/",
            self.list_pedidos,
            methods=["GET"],
        )
        self.add_api_route(
            "/",
            self.create_pedido,
            methods=["POST"],
        )
        self.add_api_route(
            "/{pedido_id}",
            self.get_pedido,
            methods=["GET"],
        )
        self.add_api_route(
            "/{pedido_id}",
            self.update_pedido,
            methods=["PUT"],
        )
        self.add_api_route(
            "/{pedido_id}",
            self.delete_pedido,
            methods=["DELETE"],
        )

    async def list_pedidos(self):

        # Logic to list pedidos goes here
        return {"message": "List of pedidos"}

    async def create_pedido(self, request: CreatePedidoRequest):
        self.service.create_pedido(request)

        # Logic to create a new pedido goes here
        return {"message": "Pedido criado"}

    async def get_pedido(self, pedido_id: UUID):
        # Logic to get a specific pedido by ID goes here
        return {"message": f"Details of pedido {pedido_id}"}

    async def update_pedido(self, pedido_id: UUID):
        # Logic to update a specific pedido by ID goes here
        return {"message": f"Pedido {pedido_id} updated"}

    async def delete_pedido(self, pedido_id: UUID):
        # Logic to delete a specific pedido by ID goes here
        return {"message": f"Pedido {pedido_id} deleted"}
