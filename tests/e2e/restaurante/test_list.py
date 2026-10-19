import pytest
from httpx import AsyncClient

from tests.e2e.restaurante.conftest import CreateRestauranteResponse


@pytest.mark.anyio
async def test_list_empty_restaurants(client: AsyncClient):
    response = await client.get("/restaurante/")

    assert response.status_code == 200
    assert response.json() == []


@pytest.mark.anyio
async def test_list_restaurants(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
):
    response = await client.get("/restaurante/")

    assert response.status_code == 200

    restaurants = response.json()

    assert len(restaurants) == 1
    assert restaurants[0]["nome"] == "Restaurante Teste"
