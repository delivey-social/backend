from src.modules.pedido.types import CreatePedidoRequest


class PedidoService:
    def __init__(self):
        pass

    def create_pedido(self, data: CreatePedidoRequest):

        print(f"Creating pedido with data: {data}")
