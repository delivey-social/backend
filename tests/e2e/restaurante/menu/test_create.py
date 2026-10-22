import pytest
from httpx import AsyncClient

from src.modules.restaurante.types import CreateMenuItemRequest
from tests.e2e.restaurante.conftest import CreateRestauranteResponse


@pytest.mark.anyio
async def test_create_menu_item(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
):
    payload = CreateMenuItemRequest(nome="Pizza", preco=300)

    response = await client.post(
        f"/menu/{create_restaurante.id}",
        json=payload.model_dump(),
    )

    assert response.status_code == 201

    res_payload = response.json()

    assert "id" in res_payload["item"]
    assert res_payload["item"]["nome"] == payload.nome
    assert res_payload["item"]["preco"] == payload.preco
