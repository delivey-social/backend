from uuid import uuid4

import pytest
from httpx import AsyncClient

from tests.e2e.restaurante.conftest import CreateRestauranteResponse


@pytest.mark.anyio
async def test_list_menu(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
):
    response = await client.get(f"/menu/{create_restaurante.id}")

    assert response.status_code == 200

    res = response.json()

    assert "id" in res
    assert "restaurante_id" in res
    assert "itens" in res


@pytest.mark.anyio
async def test_list_no_menu(client: AsyncClient):
    response = await client.get(f"/menu/{uuid4()}")

    assert response.status_code == 404
    assert response.json() == {"message": "Menu não encontrado para o restaurante"}
