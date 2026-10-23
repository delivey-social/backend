from unittest.mock import AsyncMock
from uuid import UUID

import pytest

from src.main import app
from src.modules.restaurante.domain.restaurante import Restaurante
from src.modules.restaurante.infra.dependencies import get_restaurante_service
from src.modules.restaurante.service import RestauranteService


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
