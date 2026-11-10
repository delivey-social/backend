import os

import pytest
from httpx import AsyncClient

from src.constants import KEYCLOAK_URL


@pytest.fixture()
async def auth_token() -> str:
    token_url = f"{KEYCLOAK_URL}/realms/master/protocol/openid-connect/token"

    client_id = os.getenv("E2E_CLIENT_ID")
    username = os.getenv("E2E_USERNAME")
    password = os.getenv("E2E_PASSWORD")

    async with AsyncClient() as client:
        response = await client.post(
            token_url,
            data={
                "grant_type": "password",
                "client_id": client_id,
                "username": username,
                "password": password,
            },
        )

        response.raise_for_status()
        token_data = response.json()

        return token_data["access_token"]


@pytest.fixture()
async def admin_token() -> str:
    token_url = f"{KEYCLOAK_URL}/realms/master/protocol/openid-connect/token"

    client_id = os.getenv("E2E_CLIENT_ID")
    username = os.getenv("E2E_ADMIN_USERNAME")
    password = os.getenv("E2E_ADMIN_PASSWORD")

    async with AsyncClient() as client:
        response = await client.post(
            token_url,
            data={
                "grant_type": "password",
                "client_id": client_id,
                "username": username,
                "password": password,
            },
        )

        response.raise_for_status()
        token_data = response.json()

        return token_data["access_token"]


@pytest.fixture()
async def admin_client(client: AsyncClient, admin_token: str):
    client.headers["Authorization"] = f"Bearer {admin_token}"

    yield client

    client.headers.pop("Authorization", None)
