from unittest.mock import AsyncMock
from uuid import UUID, uuid4

import pytest
from httpx import AsyncClient

from src.main import app
from src.modules.pedido.infra.dependencies import get_pedido_repo
from src.modules.pedido.infra.inmemory_repo import InMemoryPedidoRepository
from src.modules.pedido.types import CreatePedidoRequest, CreatePedidoRequestItem
from src.modules.restaurante.domain.menu import MenuItem
from src.modules.restaurante.domain.restaurante import Restaurante
from src.modules.restaurante.infra.dependencies import (
    get_menu_service,
    get_restaurante_service,
)
from src.modules.restaurante.menu_service import MenuService
from src.modules.restaurante.service import RestauranteService


@pytest.fixture(autouse=True)
async def mock_pedido_repository():
    repository = InMemoryPedidoRepository()

    async def override_get_pedido_repo():
        yield repository

    app.dependency_overrides[get_pedido_repo] = override_get_pedido_repo

    yield repository

    app.dependency_overrides.pop(get_pedido_repo, None)


@pytest.fixture
async def mock_restaurante_service():
    service = AsyncMock(spec=RestauranteService)

    async def get_restaurante(restaurant_id: UUID) -> Restaurante:
        return Restaurante(
            id=restaurant_id,
            nome="Restaurante fake",
        )

    service.get_restaurante.side_effect = get_restaurante

    async def override_get_restaurante_service():
        return service

    app.dependency_overrides[get_restaurante_service] = override_get_restaurante_service

    yield service

    app.dependency_overrides.pop(get_restaurante_service, None)


@pytest.fixture
async def mock_menu_service():
    service = AsyncMock(spec=MenuService)

    async def get_menu_item(restaurant_id: UUID, item_id: UUID) -> MenuItem:
        return MenuItem(
            id=item_id,
            nome="Item fake",
            preco=300,
        )

    service.get_menu_item.side_effect = get_menu_item

    async def override_get_menu_service():
        return service

    app.dependency_overrides[get_menu_service] = override_get_menu_service

    yield service

    app.dependency_overrides.pop(get_menu_service, None)


@pytest.fixture
async def create_pedido(
    admin_client: AsyncClient,
    mock_restaurante_service,
    mock_menu_service,
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

    return await admin_client.post(
        "/pedido/",
        json=payload.model_dump(mode="json"),
    )
