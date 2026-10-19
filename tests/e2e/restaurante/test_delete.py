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
