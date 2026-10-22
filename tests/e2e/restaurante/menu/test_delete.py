import pytest
from httpx import AsyncClient

from tests.e2e.restaurante.conftest import CreateRestauranteResponse
from tests.e2e.restaurante.menu.conftest import CreateMenuItemResponseItem


@pytest.mark.anyio
async def test_delete_menu_item(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
    create_menu_item: CreateMenuItemResponseItem,
):
    response = await client.delete(
        f"/menu/{create_restaurante.id}/{create_menu_item.id}"
    )

    assert response.status_code == 200

    data = response.json()

    assert data["message"] == "Item excluído com sucesso"
    # TODO: Assert that the item is actually deleted from the database
