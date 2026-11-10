from uuid import uuid4

import pytest
from httpx import AsyncClient, Response


@pytest.mark.anyio
async def test_delete_pedido_not_found(
    admin_client: AsyncClient,
):
    response = await admin_client.delete(f"/pedido/{uuid4()}")

    assert response.status_code == 404
    assert response.json() == {"message": "Pedido não encontrado"}


@pytest.mark.anyio
async def test_delete_pedido_success(
    admin_client: AsyncClient,
    create_pedido: Response,
):
    pedido = create_pedido.json()
    response = await admin_client.delete(f"/pedido/{pedido['id']}")

    assert response.status_code == 200
    assert response.json() == {"message": "Pedido excluído com sucesso"}


@pytest.mark.anyio
async def test_delete_pedido_twice(
    admin_client: AsyncClient,
    create_pedido: Response,
):
    pedido = create_pedido.json()
    response = await admin_client.delete(f"/pedido/{pedido['id']}")

    assert response.status_code == 200
    assert response.json() == {"message": "Pedido excluído com sucesso"}

    response = await admin_client.delete(f"/pedido/{pedido['id']}")

    assert response.status_code == 404
    assert response.json() == {"message": "Pedido não encontrado"}
