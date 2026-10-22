from uuid import UUID, uuid4

import pytest
from httpx import AsyncClient

from .conftest import CreateRestauranteResponse, GetRestauranteFixture


@pytest.mark.anyio
async def test_delete_restaurante(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
    get_restaurante: GetRestauranteFixture,
):
    response = await client.delete(f"/restaurante/{create_restaurante.id}")

    assert response.status_code == 200

    res = await get_restaurante(UUID(create_restaurante.id))

    assert res.status_code == 404


@pytest.mark.anyio
async def test_delete_restaurante_not_found(client: AsyncClient):
    response = await client.delete(f"/restaurante/{uuid4()}")

    assert response.status_code == 404
    assert response.json() == {"message": "Restaurante não encontrado"}


@pytest.mark.anyio
async def test_delete_restaurante_twice(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
):
    response = await client.delete(f"/restaurante/{create_restaurante.id}")
    assert response.status_code == 200

    # Try to delete the same restaurant again
    response = await client.delete(f"/restaurante/{create_restaurante.id}")

    assert response.status_code == 404
    assert response.json() == {"message": "Restaurante não encontrado"}
