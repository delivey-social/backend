from collections.abc import AsyncIterator, Callable
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
            service: RestauranteService = Depends(service_dependency),
        ):
            return await service.list_restaurantes()

        return endpoint

    @staticmethod
    def get_restaurante_endpoint(
        service_dependency: RestauranteServiceDependency,
    ):
        async def endpoint(
            restaurante_id: UUID,
            service: RestauranteService = Depends(service_dependency),
        ):
            return await service.get_restaurante(restaurante_id)

        return endpoint

    @staticmethod
    def create_restaurante_endpoint(
        service_dependency: RestauranteServiceDependency,
    ):
        # TODO: Create restaurante endpoint should return the created resource
        async def endpoint(
            data: CreateRestauranteRequest,
            service: RestauranteService = Depends(service_dependency),
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
            service: RestauranteService = Depends(service_dependency),
        ):
            await service.update_restaurante(
                UpdateRestauranteRequest(id=restaurante_id, nome=data.nome)
            )
            return {"message": f"Restaurante {restaurante_id} updated"}

        return endpoint

    @staticmethod
    def delete_restaurante_endpoint(
        service_dependency: RestauranteServiceDependency,
    ):
        async def endpoint(
            restaurante_id: UUID,
            service: RestauranteService = Depends(service_dependency),
        ):
            await service.delete_restaurante(restaurante_id)
            return {"message": f"Restaurante {restaurante_id} deleted"}

        return endpoint


class MenuRouter(APIRouter):
    def __init__(self, service: MenuService):
        super().__init__(prefix="/menu", tags=["Menu"])

        self.service = service

        self.add_api_route(
            "/{restaurante_id}",
            self.get_menu,
            methods=["GET"],
        )

        self.add_api_route(
            "/{restaurante_id}",
            self.create_menu_item,
            methods=["POST"],
        )

        self.add_api_route(
            "/{restaurante_id}/{item_id}",
            self.update_menu_item,
            methods=["PUT"],
        )

        self.add_api_route(
            "/{restaurante_id}/{item_id}",
            self.delete_menu_item,
            methods=["DELETE"],
        )

    async def get_menu(self, restaurante_id: UUID):
        return await self.service.get_menu(restaurante_id)

    async def create_menu_item(
        self,
        restaurante_id: UUID,
        data: CreateMenuItemRequest,
    ):
        await self.service.create_menu_item(restaurante_id, data)

        return {"message": "Menu item created successfully"}

    async def update_menu_item(
        self,
        restaurante_id: UUID,
        item_id: UUID,
        data: CreateMenuItemRequest,
    ):
        # await self.service.update_menu_item(restaurante_id, item_id, data)

        return {"message": "Menu item updated successfully"}

    async def delete_menu_item(self, restaurante_id: UUID, item_id: UUID):
        # await self.service.delete_menu_item(restaurante_id, item_id)

        return {"message": "Menu item deleted successfully"}
