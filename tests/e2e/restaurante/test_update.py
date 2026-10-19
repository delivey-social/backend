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
    assert response.json() == {
        "message": "Restaurante atualizado com sucesso",
    }
