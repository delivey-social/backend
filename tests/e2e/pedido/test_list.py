import pytest
from httpx import AsyncClient


@pytest.mark.anyio
async def test_list_pedidos_empty(client: AsyncClient):
    response = await client.get("/pedido/")

    assert response.status_code == 200

    assert response.json() == {"pedidos": []}
