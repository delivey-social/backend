from uuid import uuid4

import pytest
from httpx import AsyncClient

from tests.e2e.restaurante.conftest import CreateRestauranteResponse
from tests.e2e.restaurante.menu.conftest import CreateMenuItemResponseItem


@pytest.mark.anyio
async def test_get_menu_item(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
    create_menu_item: CreateMenuItemResponseItem,
):
    response = await client.get(
        f"/menu/{create_restaurante.id}/{create_menu_item.id}",
    )

    assert response.status_code == 200

    assert response.json() == {
        "id": create_menu_item.id,
        "nome": create_menu_item.nome,
        "preco": create_menu_item.preco,
    }


@pytest.mark.anyio
async def test_get_menu_item_not_found(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
):
    response = await client.get(
        f"/menu/{create_restaurante.id}/{uuid4()}",
    )

    assert response.status_code == 404

    assert response.json() == {"message": "Item não encontrado"}
