import pytest
from httpx import AsyncClient
from pydantic import BaseModel

from src.modules.restaurante.types import CreateRestauranteRequest


class CreateRestauranteResponse(BaseModel):
    message: str = "Restaurante criado com sucesso!"


@pytest.fixture
async def create_restaurante(
    client: AsyncClient,
) -> CreateRestauranteResponse:
    """Fixture that creates a restaurant and returns the payload/data."""
    payload = CreateRestauranteRequest(
        nome="Restaurante Teste",
    )

    response = await client.post("/restaurante/", json=payload.model_dump())

    assert response.status_code == 201

    return CreateRestauranteResponse()
