from uuid import uuid4

import pytest
from httpx import AsyncClient

from tests.e2e.restaurante.conftest import CreateRestauranteResponse


@pytest.mark.anyio
async def test_get_logo_restaurante_not_found(
    client: AsyncClient,
):
    non_existent_id = uuid4()

    res = await client.get(f"/restaurante/{non_existent_id}/logo")

    assert res.status_code == 404


@pytest.mark.anyio
async def test_get_logo_invalid_restaurante_id(
    client: AsyncClient,
):
    invalid_id = "invalid-uuid"

    res = await client.get(f"/restaurante/{invalid_id}/logo")

    assert res.status_code == 422


@pytest.mark.anyio
async def test_get_logo_no_logo(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
):
    data = create_restaurante.model_dump()
    id = data.get("id")

    res = await client.get(f"/restaurante/{id}/logo")

    assert res.status_code == 404


@pytest.mark.anyio
async def test_get_logo_success(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
    tmp_path: str,
):
    data = create_restaurante.model_dump()
    id = data.get("id")

    file = b"fake image data"  # Simulating a fake image file

    res = await client.put(
        f"/restaurante/{id}/logo",
        files={"file": ("logo.png", file, "image/png")},
    )

    assert res.status_code == 200

    res = await client.get(f"/restaurante/{id}/logo")

    assert res.status_code == 200
    assert res.headers["content-type"] == "image/*"
