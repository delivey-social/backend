import pytest
from httpx import AsyncClient
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
