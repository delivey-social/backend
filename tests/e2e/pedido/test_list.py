import pytest
from httpx import AsyncClient, Response


@pytest.mark.anyio
async def test_list_pedidos_empty(client: AsyncClient):
    response = await client.get("/pedido/")

    assert response.status_code == 200

    assert response.json() == {"pedidos": []}


@pytest.mark.anyio
async def test_list_pedidos(
    client: AsyncClient,
    create_pedido: Response,
):
    response = await client.get("/pedido/")

    assert response.status_code == 200

    json = response.json()

    assert "pedidos" in json
    assert len(json["pedidos"]) == 1
