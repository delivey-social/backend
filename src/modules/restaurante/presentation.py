from collections.abc import AsyncIterator, Callable
from typing import Annotated
from uuid import UUID

from fastapi import APIRouter, Depends
from fastapi.responses import JSONResponse

from src.modules.restaurante.menu_service import MenuService
from src.modules.restaurante.types import (
    CreateMenuItemRequest,
    CreateRestauranteRequest,
    UpdateRestauranteRequest,
)

from .service import RestauranteService

RestauranteServiceDependency = Callable[..., AsyncIterator[RestauranteService]]
MenuServiceDependency = Callable[..., AsyncIterator[MenuService]]


class RestauranteRouter(APIRouter):
    def __init__(
        self,
        service_dependency: RestauranteServiceDependency,
    ):
        super().__init__(prefix="/restaurante", tags=["Restaurante"])

        self.add_api_route(
            "/",
            self.list_restaurantes_endpoint(service_dependency),
            methods=["GET"],
        )
        self.add_api_route(
            "/{restaurante_id}",
            self.get_restaurante_endpoint(service_dependency),
            methods=["GET"],
        )

        self.add_api_route(
            "/",
            self.create_restaurante_endpoint(service_dependency),
            methods=["POST"],
        )

        self.add_api_route(
            "/{restaurante_id}",
            self.update_restaurante_endpoint(service_dependency),
            methods=["PUT"],
        )

        self.add_api_route(
            "/{restaurante_id}",
            self.delete_restaurante_endpoint(service_dependency),
            methods=["DELETE"],
        )

    @staticmethod
    def list_restaurantes_endpoint(
        service_dependency: RestauranteServiceDependency,
    ):
        async def endpoint(
            service: Annotated[RestauranteService, Depends(service_dependency)],
        ):
            return await service.list_restaurantes()

        return endpoint

    @staticmethod
    def get_restaurante_endpoint(
        service_dependency: RestauranteServiceDependency,
    ):
        async def endpoint(
            restaurante_id: UUID,
            service: Annotated[RestauranteService, Depends(service_dependency)],
        ):
            return await service.get_restaurante(restaurante_id)

        return endpoint

    @staticmethod
    def create_restaurante_endpoint(
        service_dependency: RestauranteServiceDependency,
    ):
        async def endpoint(
            data: CreateRestauranteRequest,
            service: Annotated[RestauranteService, Depends(service_dependency)],
        ):
            restaurante = await service.create_restaurante(data)

            return JSONResponse(
                status_code=201,
                content=restaurante.model_dump(mode="json"),
            )

        return endpoint

    @staticmethod
    def update_restaurante_endpoint(
        service_dependency: RestauranteServiceDependency,
    ):
        async def endpoint(
            restaurante_id: UUID,
            data: CreateRestauranteRequest,
            service: Annotated[RestauranteService, Depends(service_dependency)],
        ):
            restaurante = await service.update_restaurante(
                UpdateRestauranteRequest(id=restaurante_id, nome=data.nome)
            )

            return JSONResponse(
                status_code=200,
                content=restaurante.model_dump(mode="json"),
            )

        return endpoint

    @staticmethod
    def delete_restaurante_endpoint(
        service_dependency: RestauranteServiceDependency,
    ):
        async def endpoint(
            restaurante_id: UUID,
            service: Annotated[RestauranteService, Depends(service_dependency)],
        ):
            await service.delete_restaurante(restaurante_id)

            return JSONResponse(
                status_code=200,
                content={"message": "Restaurante excluído com sucesso"},
            )

        return endpoint


class MenuRouter(APIRouter):
    def __init__(self, service: MenuServiceDependency):
        super().__init__(prefix="/menu", tags=["Menu"])

        self.add_api_route(
            "/{restaurante_id}",
            self.get_menu_endpoint(service),
            methods=["GET"],
        )

        self.add_api_route(
            "/{restaurante_id}",
            self.create_menu_item_endpoint(service),
            methods=["POST"],
        )

        self.add_api_route(
            "/{restaurante_id}/{item_id}",
            self.update_menu_item_endpoint(service),
            methods=["PUT"],
        )

        self.add_api_route(
            "/{restaurante_id}/{item_id}",
            self.delete_menu_item_endpoint(service),
            methods=["DELETE"],
        )

        self.add_api_route(
            "/{restaurante_id}/{item_id}",
            self.get_menu_item_endpoint(service),
            methods=["GET"],
        )

    @staticmethod
    def get_menu_endpoint(service_dependency: MenuServiceDependency):
        async def endpoint(
            restaurante_id: UUID,
            service: Annotated[MenuService, Depends(service_dependency)],
        ):
            return await service.get_menu(restaurante_id)

        return endpoint

    @staticmethod
    def create_menu_item_endpoint(service_dependency: MenuServiceDependency):
        async def endpoint(
            restaurante_id: UUID,
            data: CreateMenuItemRequest,
            service: Annotated[MenuService, Depends(service_dependency)],
        ):
            item = await service.create_menu_item(restaurante_id, data)

            return JSONResponse(
                status_code=201,
                content={
                    "message": "Menu item created successfully",
                    "item": item.model_dump(mode="json"),
                },
            )

        return endpoint

    @staticmethod
    def update_menu_item_endpoint(service_dependency: MenuServiceDependency):
        async def endpoint(
            restaurante_id: UUID,
            item_id: UUID,
            data: CreateMenuItemRequest,
            service: Annotated[MenuService, Depends(service_dependency)],
        ):
            item = await service.update_menu_item(restaurante_id, item_id, data)

            return JSONResponse(
                status_code=200,
                content={
                    "message": "Item atualizado com sucesso",
                    "item": item.model_dump(mode="json"),
                },
            )

        return endpoint

    @staticmethod
    def delete_menu_item_endpoint(service_dependency: MenuServiceDependency):
        async def endpoint(
            restaurante_id: UUID,
            item_id: UUID,
            service: Annotated[MenuService, Depends(service_dependency)],
        ):
            await service.delete_menu_item(restaurante_id, item_id)

            return JSONResponse(
                status_code=200,
                content={"message": "Item excluído com sucesso"},
            )

        return endpoint

    @staticmethod
    def get_menu_item_endpoint(service_dependency: MenuServiceDependency):
        async def endpoint(
            restaurante_id: UUID,
            item_id: UUID,
            service: Annotated[MenuService, Depends(service_dependency)],
        ):
            return await service.get_menu_item(restaurante_id, item_id)

        return endpoint
