from uuid import uuid4

import pytest
from httpx import AsyncClient


@pytest.mark.anyio
async def test_update_not_found(client: AsyncClient):
    response = await client.put(f"/pedido/{uuid4()}")

    assert response.status_code == 404
