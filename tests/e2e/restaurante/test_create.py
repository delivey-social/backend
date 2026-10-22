from uuid import UUID

import pytest

from .conftest import CreateRestauranteResponse, GetRestauranteFixture


@pytest.mark.anyio
async def test_create_restaurante(
    create_restaurante: CreateRestauranteResponse,
    get_restaurante: GetRestauranteFixture,
):
    data = create_restaurante.model_dump()

    assert "id" in data
    assert "nome" in data

    response = await get_restaurante(UUID(data["id"]))

    assert response.status_code == 200

    assert response.data.id == data["id"]
    assert response.data.nome == data["nome"]
