from uuid import uuid4

import pytest
from httpx import AsyncClient, Response


@pytest.mark.anyio
async def test_update(
    client: AsyncClient,
    create_pedido: Response,
):
    pedido = create_pedido.json()
    pedido_id = pedido["id"]

    response = await client.put(
        f"/pedido/{pedido_id}",
    )

    assert response.status_code == 200


@pytest.mark.anyio
async def test_update_not_found(client: AsyncClient):
    response = await client.put(f"/pedido/{uuid4()}")

    assert response.status_code == 404
