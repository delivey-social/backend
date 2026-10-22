from collections.abc import Awaitable, Callable
from uuid import UUID

import pytest
from httpx import AsyncClient, Response
from pydantic import BaseModel

from src.modules.restaurante.types import CreateRestauranteRequest


class GetRestauranteResponse(BaseModel):
    id: str
    nome: str


class CreateRestauranteResponse(BaseModel):
    id: str
    nome: str


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

    return CreateRestauranteResponse(**response.json())


class GetRestauranteResponseWrapper:
    def __init__(self, response: Response):
        self._response = response

    @property
    def data(self) -> GetRestauranteResponse:
        """Parses and returns the typed model only when accessed."""
        return GetRestauranteResponse(**self._response.json())

    def __getattr__(self, name: str):
        return getattr(self._response, name)


GetRestauranteFixture = Callable[
    [UUID],
    Awaitable[GetRestauranteResponseWrapper],
]


@pytest.fixture
async def get_restaurante(client: AsyncClient):
    async def _get_restaurante(restaurant_id: UUID):
        response = await client.get(f"/restaurante/{restaurant_id}")

        return GetRestauranteResponseWrapper(response)

    return _get_restaurante
