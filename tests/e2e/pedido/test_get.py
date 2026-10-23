from uuid import uuid4

import pytest
from httpx import AsyncClient, Response


@pytest.mark.anyio
async def test_get_pedido_not_found(
    client: AsyncClient,
):
    response = await client.get(f"/pedido/{uuid4()}")

    assert response.status_code == 404


@pytest.mark.anyio
async def test_get_pedido_success(
    client: AsyncClient,
    create_pedido: Response,
):
    data = create_pedido.json()

    pedido_id = data["id"]
    response = await client.get(f"/pedido/{pedido_id}")

    assert response.status_code == 200

    data = response.json()
    assert data["message"] == f"Details of pedido {pedido_id}"
