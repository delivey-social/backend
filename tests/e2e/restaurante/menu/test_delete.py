from uuid import UUID, uuid4

import pytest
from httpx import AsyncClient

from tests.e2e.restaurante.conftest import CreateRestauranteResponse
from tests.e2e.restaurante.menu.conftest import (
    CreateMenuItemResponseItem,
    GetMenuItemFixture,
)


@pytest.mark.anyio
async def test_delete_menu_item(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
    create_menu_item: CreateMenuItemResponseItem,
    get_menu_item: GetMenuItemFixture,
):
    response = await client.delete(
        f"/menu/{create_restaurante.id}/{create_menu_item.id}"
    )

    assert response.status_code == 200

    data = response.json()

    assert data["message"] == "Item excluído com sucesso"

    res = await get_menu_item(UUID(create_menu_item.id))

    assert res.status_code == 404


@pytest.mark.anyio
async def test_delete_menu_item_not_found(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
):
    response = await client.delete(f"/menu/{create_restaurante.id}/{uuid4()}")

    assert response.status_code == 404

    data = response.json()

    assert data["message"] == "Item não encontrado"


@pytest.mark.anyio
async def test_delete_menu_item_twice(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
    create_menu_item: CreateMenuItemResponseItem,
):
    response = await client.delete(
        f"/menu/{create_restaurante.id}/{create_menu_item.id}"
    )
    assert response.status_code == 200

    response = await client.delete(
        f"/menu/{create_restaurante.id}/{create_menu_item.id}"
    )

    assert response.status_code == 404

    data = response.json()
    assert data["message"] == "Item não encontrado"
