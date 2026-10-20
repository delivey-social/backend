from uuid import uuid4

import pytest
from httpx import AsyncClient


@pytest.mark.anyio
async def test_list_no_menu(client: AsyncClient):
    response = await client.get(f"/menu/{uuid4()}")

    assert response.status_code == 404
    assert response.json() == {"message": "Menu não encontrado para o restaurante"}
