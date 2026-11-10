import pytest
from httpx import AsyncClient


@pytest.mark.anyio
async def test_ping_admin(admin_client: AsyncClient):
    response = await admin_client.get("/admin")

    assert response.status_code == 200


@pytest.mark.anyio
async def test_ping_admin_without_token(client: AsyncClient):
    response = await client.get("/admin")

    assert response.status_code == 401


@pytest.mark.anyio
async def test_ping_admin_with_invalid_token(
    auth_token: str,
    client: AsyncClient,
):
    response = await client.get(
        "/admin", headers={"Authorization": f"Bearer {auth_token}"}
    )

    assert response.status_code
