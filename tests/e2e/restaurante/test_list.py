import pytest
from httpx import AsyncClient


@pytest.mark.anyio
async def test_list_empty_restaurants(client: AsyncClient):
    response = await client.get("/restaurante/")

    assert response.status_code == 200
    assert response.json() == []
