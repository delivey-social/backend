import tomllib

from src.constants import ROOT_PATH


def get_app_version() -> str:
    pyproject_path = ROOT_PATH / "pyproject.toml"

    if not pyproject_path.exists():
        raise FileNotFoundError(f"File not found: {pyproject_path}")

    with open(pyproject_path, "rb") as f:
        data = tomllib.load(f)

        version = data.get("project", {}).get("version")

        if not version:
            raise ValueError("Version not found in pyproject.toml")

    return version
