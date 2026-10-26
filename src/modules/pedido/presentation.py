from collections.abc import Awaitable, Callable
from uuid import UUID

from fastapi import APIRouter, Depends
from fastapi.responses import JSONResponse

from src.modules.pedido.service import PedidoService
from src.modules.pedido.types import CreatePedidoRequest

PedidoServiceDependency = Callable[..., Awaitable[PedidoService]]


class PedidoRouter(APIRouter):
    def __init__(self, service_dependency: PedidoServiceDependency):
        super().__init__(prefix="/pedido", tags=["Pedido"])

        self.add_api_route(
            "/",
            self.list_pedidos_endpoint(service_dependency),
            methods=["GET"],
        )
        self.add_api_route(
            "/{pedido_id}",
            self.get_pedido_endpoint(service_dependency),
            methods=["GET"],
        )
        self.add_api_route(
            "/",
            self.create_pedido_endpoint(service_dependency),
            methods=["POST"],
        )
        self.add_api_route(
            "/{pedido_id}",
            self.update_pedido_endpoint(service_dependency),
            methods=["PUT"],
        )
        self.add_api_route(
            "/{pedido_id}",
            self.delete_pedido_endpoint(service_dependency),
            methods=["DELETE"],
        )

    @staticmethod
    def list_pedidos_endpoint(service_dependency: PedidoServiceDependency):
        async def endpoint(
            service: PedidoService = Depends(service_dependency),
        ):
            pedidos = await service.list_pedidos()

            return JSONResponse(
                content={
                    "pedidos": [pedido.model_dump(mode="json") for pedido in pedidos]
                }
            )

        return endpoint

    @staticmethod
    def get_pedido_endpoint(service_dependency: PedidoServiceDependency):
        async def endpoint(
            pedido_id: UUID,
            service: PedidoService = Depends(service_dependency),
        ):
            pedido = await service.get_pedido(pedido_id)

            return JSONResponse(
                content={**pedido.model_dump(mode="json")}, status_code=200
            )

        return endpoint

    @staticmethod
    def create_pedido_endpoint(service_dependency: PedidoServiceDependency):
        async def endpoint(
            request: CreatePedidoRequest,
            service: PedidoService = Depends(service_dependency),
        ):
            pedido = await service.create_pedido(request)

            return JSONResponse(
                content={**pedido.model_dump(mode="json")}, status_code=201
            )

        return endpoint

    @staticmethod
    def update_pedido_endpoint(service_dependency: PedidoServiceDependency):
        async def endpoint(
            pedido_id: UUID,
            service: PedidoService = Depends(service_dependency),
        ):
            await service.update_pedido(pedido_id)

            return {"message": f"Pedido {pedido_id} updated"}

        return endpoint

    @staticmethod
    def delete_pedido_endpoint(service_dependency: PedidoServiceDependency):
        async def endpoint(
            pedido_id: UUID,
            service: PedidoService = Depends(service_dependency),
        ):
            await service.delete_pedido(pedido_id)

            return {"message": "Pedido excluído com sucesso"}

        return endpoint
