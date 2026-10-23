from uuid import uuid4

import pytest
from httpx import AsyncClient


@pytest.mark.anyio
async def test_get_pedido(
    client: AsyncClient,
):
    response = await client.get(f"/pedido/{uuid4()}")

    assert response.status_code == 200
