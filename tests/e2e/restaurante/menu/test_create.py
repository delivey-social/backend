from uuid import uuid4

import pytest
from httpx import AsyncClient

from src.modules.restaurante.types import CreateMenuItemRequest
from tests.e2e.restaurante.menu.conftest import CreateMenuItemResponseItem


@pytest.mark.anyio
async def test_create_menu_item(create_menu_item: CreateMenuItemResponseItem):
    data = create_menu_item.model_dump()

    assert "id" in data
    assert "nome" in data
    assert "preco" in data


@pytest.mark.anyio
async def test_create_menu_item_no_restaurante(client: AsyncClient):
    payload = CreateMenuItemRequest(nome="Pizza", preco=300)

    response = await client.post(
        f"/menu/{uuid4()}",
        json=payload.model_dump(),
    )

    assert response.status_code == 404
    assert response.json() == {
        "message": "Menu não encontrado para o restaurante",
    }

    # TODO: Assert that the item is actually created in the database
