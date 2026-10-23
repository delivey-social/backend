from uuid import uuid4

import pytest
from httpx import AsyncClient

from src.modules.pedido.types import CreatePedidoRequest, CreatePedidoRequestItem


@pytest.mark.anyio
async def test_create_pedido(
    client: AsyncClient,
    mock_restaurante_service,
):
    restaurant_id = uuid4()

    payload = CreatePedidoRequest(
        user_id=uuid4(),
        restaurant_id=restaurant_id,
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

    assert response.status_code == 201
    mock_restaurante_service.get_restaurante.assert_awaited_once_with(restaurant_id)
