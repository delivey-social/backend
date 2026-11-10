import pytest
from httpx import AsyncClient


@pytest.mark.anyio
async def test_no_token(client: AsyncClient):
    response = await client.get("/secured")

    assert response.status_code == 401


@pytest.mark.anyio
async def test_invalid_token(client: AsyncClient):
    response = await client.get(
        "/secured",
        headers={
            "Authorization": "Bearer invalid_token",
        },
    )

    assert response.status_code == 401


@pytest.mark.anyio
async def test_valid_token(client: AsyncClient, auth_token: str):
    response = await client.get(
        "/secured",
        headers={
            "Authorization": f"Bearer {auth_token}",
        },
    )

    assert response.status_code == 200
