from pathlib import Path
from uuid import uuid4

import pytest
from httpx import AsyncClient

from tests.e2e.restaurante.conftest import CreateRestauranteResponse


@pytest.mark.anyio
async def test_upload_logo(
    client: AsyncClient, create_restaurante: CreateRestauranteResponse, tmp_path: str
):
    data = create_restaurante.model_dump()
    id = data.get("id")

    file = b"fake image data"  # Simulating a fake image file

    res = await client.put(
        f"/restaurante/{id}/logo",
        files={"file": ("logo.png", file, "image/png")},
    )

    print(res.json())

    assert res.status_code == 200

    expected_file_path = Path(tmp_path) / "logos" / f"{id}.png"

    assert expected_file_path.exists()
    assert expected_file_path.read_bytes() == file


@pytest.mark.anyio
async def test_upload_logo_multiple_files(
    client: AsyncClient, create_restaurante: CreateRestauranteResponse, tmp_path: str
):
    data = create_restaurante.model_dump()
    id = data.get("id")

    file1 = b"fake image data 1"
    file2 = b"fake image data 2"

    res = await client.put(
        f"/restaurante/{id}/logo",
        files=[
            ("file", ("logo1.png", file1, "image/png")),
            ("file", ("logo2.png", file2, "image/png")),
        ],
    )

    assert res.status_code == 400
    assert res.json() == {
        "message": "Apenas um arquivo de imagem deve ser enviado",
    }

    logo_paths = Path(tmp_path) / "logos"

    assert not logo_paths.exists() or not any(logo_paths.iterdir())


@pytest.mark.anyio
async def test_upload_logo_no_file(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
    tmp_path: str,
):
    data = create_restaurante.model_dump()
    id = data.get("id")

    res = await client.put(
        f"/restaurante/{id}/logo",
        files={},
    )

    assert res.status_code == 422

    logo_paths = Path(tmp_path) / "logos"

    assert not logo_paths.exists() or not any(logo_paths.iterdir())


@pytest.mark.anyio
async def test_upload_logo_invalid_restaurante_id(
    client: AsyncClient,
    tmp_path: str,
):
    invalid_id = "invalid-uuid"

    file = b"fake image data"

    res = await client.put(
        f"/restaurante/{invalid_id}/logo",
        files={"file": ("logo.png", file, "image/png")},
    )

    assert res.status_code == 422

    logo_paths = Path(tmp_path) / "logos"

    assert not logo_paths.exists() or not any(logo_paths.iterdir())


@pytest.mark.anyio
async def test_upload_logo_restaurante_not_found(
    client: AsyncClient,
    tmp_path: str,
):
    non_existent_id = uuid4()

    file = b"fake image data"

    res = await client.put(
        f"/restaurante/{non_existent_id}/logo",
        files={"file": ("logo.png", file, "image/png")},
    )

    assert res.status_code == 404
    assert res.json() == {
        "message": "Restaurante não encontrado",
    }

    logo_paths = Path(tmp_path) / "logos"

    assert not logo_paths.exists() or not any(logo_paths.iterdir())


@pytest.mark.anyio
async def test_upload_logo_unsupported_format(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
    tmp_path: str,
):
    data = create_restaurante.model_dump()
    id = data.get("id")

    file = b"fake image data"

    res = await client.put(
        f"/restaurante/{id}/logo",
        files={"file": ("logo.txt", file, "text/plain")},
    )

    assert res.status_code == 400
    assert res.json() == {
        "message": "Formato de imagem não suportado",
    }

    logo_paths = Path(tmp_path) / "logos"

    assert not logo_paths.exists() or not any(logo_paths.iterdir())


@pytest.mark.anyio
async def test_upload_logo_empty_file(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
    tmp_path: str,
):
    data = create_restaurante.model_dump()
    id = data.get("id")

    empty_file = b""

    res = await client.put(
        f"/restaurante/{id}/logo",
        files={"file": ("logo.png", empty_file, "image/png")},
    )

    assert res.status_code == 400
    assert res.json() == {
        "message": "Arquivo de imagem vazio",
    }

    logo_paths = Path(tmp_path) / "logos"

    assert not logo_paths.exists() or not any(logo_paths.iterdir())


@pytest.mark.anyio
async def test_upload_logo_too_large_file(
    client: AsyncClient,
    create_restaurante: CreateRestauranteResponse,
    tmp_path: str,
):
    data = create_restaurante.model_dump()
    id = data.get("id")

    # Simulating a large file (e.g., 10MB)
    large_file = b"a" * (10 * 1024 * 1024)

    res = await client.put(
        f"/restaurante/{id}/logo",
        files={"file": ("logo.png", large_file, "image/png")},
    )

    assert res.status_code == 400
    assert res.json() == {
        "message": "Arquivo de imagem muito grande (máximo 5MB)",
    }

    logo_paths = Path(tmp_path) / "logos"

    assert not logo_paths.exists() or not any(logo_paths.iterdir())
