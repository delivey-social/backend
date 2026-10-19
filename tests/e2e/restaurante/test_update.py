import pytest
from httpx import AsyncClient

from tests.e2e.restaurante.conftest import CreateRestauranteResponse


@pytest.mark.anyio
async def test_update_restaurante(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
):
    response = await client.put(
        f"/restaurante/{create_restaurante.id}",
        json={"nome": "Restaurante Atualizado"},
    )

    assert response.status_code == 200

    response_data = response.json()

    assert "id" in response_data
    assert "nome" in response_data
