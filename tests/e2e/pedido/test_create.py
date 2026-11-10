from unittest.mock import AsyncMock
from uuid import uuid4

import pytest
from httpx import Response

from src.modules.pedido.types import CreatePedidoRequest, CreatePedidoRequestItem


@pytest.mark.anyio
async def test_create_pedido(
    create_pedido: Response,
    mock_restaurante_service: AsyncMock,
):
    assert create_pedido.status_code == 201

    mock_restaurante_service.get_restaurante.assert_called_once()

    data = create_pedido.json()

    assert "id" in data


@pytest.mark.anyio
async def test_create_pedido_not_admin(
    client: AsyncMock,
    mock_restaurante_service: AsyncMock,
):
    payload = CreatePedidoRequest(
        user_id=uuid4(),
        restaurant_id=uuid4(),
        items=[
            CreatePedidoRequestItem(
                product_id=uuid4(),
                quantity=1,
            )
        ],
    )

    response = await client.post(
        "/pedido/",
        json=payload.model_dump(mode="json"),
    )

    assert response.status_code == 401

    mock_restaurante_service.get_restaurante.assert_not_called()
