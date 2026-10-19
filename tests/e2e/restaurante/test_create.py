import pytest

from .conftest import CreateRestauranteResponse


@pytest.mark.anyio
async def test_create_restaurante(
    create_restaurante: CreateRestauranteResponse,
):
    data = create_restaurante.model_dump()

    assert "id" in data
    assert "nome" in data
