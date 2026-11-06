import pytest
from httpx import AsyncClient

from tests.e2e.restaurante.conftest import CreateRestauranteResponse


@pytest.mark.anyio
async def test_upload_logo(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
):
    data = create_restaurante.model_dump()
    id = data.get("id")

    file = b"fake image data"  # Simulating a fake image file

    res = await client.put(
        f"/restaurante/{id}/logo",
        files={"file": ("logo.png", file, "image/png")},
    )

    print(res.json())

    assert res.status_code == 200
