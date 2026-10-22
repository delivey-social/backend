import pytest
from httpx import AsyncClient

from src.modules.restaurante.types import UpdateMenuItemRequest
from tests.e2e.restaurante.conftest import CreateRestauranteResponse
from tests.e2e.restaurante.menu.conftest import CreateMenuItemResponseItem


@pytest.mark.anyio
async def test_update_menu_item(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
    create_menu_item: CreateMenuItemResponseItem,
):
    payload = UpdateMenuItemRequest(nome="Pizza", preco=300)

    response = await client.put(
        f"/menu/{create_restaurante.id}/{create_menu_item.id}",
        json=payload.model_dump(),
    )

    assert response.status_code == 200

    data = response.json()
    item = data["item"]

    assert item["id"] == create_menu_item.id
    assert item["nome"] == payload.nome
    assert item["preco"] == payload.preco
