import pytest
from httpx import AsyncClient


@pytest.mark.anyio
async def test_ping(client: AsyncClient):
    response = await client.get("/")

    assert response.status_code == 200
    assert response.json() == {"message": "comida.app backend is alive!"}
