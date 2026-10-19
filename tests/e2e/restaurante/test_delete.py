from uuid import uuid4

import pytest
from httpx import AsyncClient

from tests.e2e.restaurante.conftest import CreateRestauranteResponse


@pytest.mark.anyio
async def test_delete_restaurante(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
):
    response = await client.delete(f"/restaurante/{create_restaurante.id}")

    assert response.status_code == 204


@pytest.mark.anyio
async def test_delete_restaurante_not_found(client: AsyncClient):
    response = await client.delete(f"/restaurante/{uuid4()}")

    assert response.status_code == 404
    assert response.json() == {"message": "Restaurante não encontrado"}
