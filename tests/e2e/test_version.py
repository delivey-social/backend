import pytest
from httpx import AsyncClient


@pytest.mark.anyio
async def test_version(client: AsyncClient):
    response = await client.get("/version")

    assert response.status_code == 200
    assert response.json() == {"version": "0.0.0"}
