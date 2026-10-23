from uuid import uuid4

import pytest
from httpx import AsyncClient


@pytest.mark.anyio
async def test_delete_pedido_not_found(
    client: AsyncClient,
):
    response = await client.delete(f"/pedido/{uuid4()}")

    assert response.status_code == 404
    assert response.json() == {"message": "Pedido não encontrado"}
