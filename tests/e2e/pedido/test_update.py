from uuid import uuid4

import pytest
from httpx import AsyncClient, Response

from src.modules.pedido.types import CreatePedidoRequestItem, UpdatePedidoRequest

payload = UpdatePedidoRequest(
    user_id=uuid4(),
    restaurant_id=uuid4(),
    items=[
        CreatePedidoRequestItem(
            product_id=uuid4(),
            quantity=2,
        )
    ],
)


@pytest.mark.anyio
async def test_update(
    client: AsyncClient,
    create_pedido: Response,
):
    pedido = create_pedido.json()
    pedido_id = pedido["id"]

    response = await client.put(
        f"/pedido/{pedido_id}",
        json=payload.model_dump(mode="json"),
    )

    assert response.status_code == 200


@pytest.mark.anyio
async def test_update_not_found(client: AsyncClient):
    response = await client.put(
        f"/pedido/{uuid4()}",
        json=payload.model_dump(mode="json"),
    )

    assert response.status_code == 404


@pytest.mark.anyio
async def test_update_result(client: AsyncClient, create_pedido: Response):
    pedido = create_pedido.json()
    pedido_id = pedido["id"]

    response = await client.put(
        f"/pedido/{pedido_id}",
        json=payload.model_dump(mode="json"),
    )

    assert response.status_code == 200
    assert response.json() == {"message": "Pedido atualizado com sucesso"}

    await client.get(f"/pedido/{pedido_id}")
