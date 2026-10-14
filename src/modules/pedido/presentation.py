from fastapi import APIRouter


class PedidoRouter(APIRouter):
    def __init__(self):
        super().__init__(prefix="/pedido", tags=["Pedido"])

        self.add_api_route("/", self.list_pedidos, methods=["GET"])

    async def list_pedidos(self):
        # Logic to list pedidos goes here
        return {"message": "List of pedidos"}
