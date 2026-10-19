import pytest
from httpx import AsyncClient

from src.modules.restaurante.types import CreateRestauranteRequest


@pytest.mark.anyio
async def test_create_restaurante(client: AsyncClient):
    payload = CreateRestauranteRequest(
        nome="Restaurante Teste",
    )

    response = await client.post(
        "/restaurante/",
        json=payload.model_dump(),
    )

    assert response.status_code == 201
    assert response.json() == {
        "message": "Restaurante criado com sucesso!",
    }
