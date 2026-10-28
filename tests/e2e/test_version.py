import re

import pytest
from httpx import AsyncClient

from src.core.scripts.get_app_version import get_app_version


@pytest.mark.anyio
async def test_version(client: AsyncClient):
    response = await client.get("/version")

    assert response.status_code == 200

    json = response.json()

    assert "version" in json

    semver_pattern = r"^\d+.\d+\.\d+$"

    version = json["version"]

    assert isinstance(version, str)
    assert re.match(semver_pattern, version), (
        f"Version '{version}' does not match semantic versioning pattern"
    )

    assert version == get_app_version()
