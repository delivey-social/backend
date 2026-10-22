from collections.abc import Awaitable, Callable
from uuid import UUID

import pytest
from httpx import AsyncClient, Response
from pydantic import BaseModel

from src.modules.restaurante.types import CreateMenuItemRequest
from tests.e2e.restaurante.conftest import CreateRestauranteResponse


class CreateMenuItemResponseItem(BaseModel):
    id: str
    nome: str
    preco: int


class CreateMenuItemResponse(BaseModel):
    message: str
    item: CreateMenuItemResponseItem


@pytest.fixture
async def create_menu_item(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
) -> CreateMenuItemResponseItem:
    payload = CreateMenuItemRequest(nome="Pizza", preco=300)

    response = await client.post(
        f"/menu/{create_restaurante.id}",
        json=payload.model_dump(),
    )

    assert response.status_code == 201

    data = response.json()

    return CreateMenuItemResponseItem(**data["item"])


class GetMenuItemResponseWrapper:
    def __init__(self, response: Response):
        self._response = response

    @property
    def data(self) -> CreateMenuItemResponseItem:
        """Parses and returns the typed model only when accessed."""
        return CreateMenuItemResponseItem(**self._response.json())

    def __getattr__(self, name: str):
        return getattr(self._response, name)


GetMenuItemFixture = Callable[[UUID], Awaitable[GetMenuItemResponseWrapper]]


@pytest.fixture
async def get_menu_item(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
):
    async def _get_menu_item(item_id: UUID) -> GetMenuItemResponseWrapper:
        response = await client.get(f"/menu/{create_restaurante.id}/{item_id}")

        return GetMenuItemResponseWrapper(response)

    return _get_menu_item
