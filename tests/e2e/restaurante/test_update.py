from uuid import UUID, uuid4

import pytest
from httpx import AsyncClient

from tests.e2e.restaurante.conftest import (
    CreateRestauranteResponse,
    GetRestauranteFixture,
)


@pytest.mark.anyio
async def test_update_restaurante(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
    get_restaurante: GetRestauranteFixture,
):
    response = await client.put(
        f"/restaurante/{create_restaurante.id}",
        json={"nome": "Restaurante Atualizado"},
    )

    assert response.status_code == 200

    data = response.json()

    assert "id" in data
    assert "nome" in data

    response = await get_restaurante(UUID(data["id"]))

    assert response.status_code == 200

    assert response.data.id == data["id"]
    assert response.data.nome == data["nome"]


@pytest.mark.anyio
async def test_update_restaurante_not_found(client: AsyncClient):
    response = await client.put(
        f"/restaurante/{uuid4()}",
        json={"nome": "Restaurante Atualizado"},
    )

    assert response.status_code == 404
    assert response.json() == {"message": "Restaurante não encontrado"}
