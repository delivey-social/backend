from unittest.mock import AsyncMock

import pytest
from httpx import AsyncClient, Response


@pytest.mark.anyio
async def test_create_pedido(
    client: AsyncClient,
    create_pedido: Response,
    mock_restaurante_service: AsyncMock,
):
    assert create_pedido.status_code == 201

    mock_restaurante_service.get_restaurante.assert_called_once()

    data = create_pedido.json()

    assert "id" in data
