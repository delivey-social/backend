from uuid import uuid4

import pytest
from httpx import AsyncClient

from .conftest import CreateRestauranteResponse


@pytest.mark.anyio
async def test_restaurant_details(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
):
    restaurant_id = create_restaurante.id
    response = await client.get(f"/restaurante/{restaurant_id}")

    assert response.status_code == 200
    data = response.json()

    assert data["id"] == restaurant_id
    assert data["nome"] == create_restaurante.nome


@pytest.mark.anyio
async def test_restaurant_details_not_found(client: AsyncClient):
    response = await client.get(f"/restaurante/{uuid4()}")

    assert response.status_code == 404
    assert response.json() == {"message": "Restaurante não encontrado"}
