import pytest

from .conftest import CreateRestauranteResponse


@pytest.mark.anyio
async def test_create_restaurante(
    create_restaurante: CreateRestauranteResponse,
):
    assert create_restaurante.model_dump() == {
        "message": "Restaurante criado com sucesso!",
    }
