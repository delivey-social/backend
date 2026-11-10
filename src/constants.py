import os
from pathlib import Path

from dotenv import load_dotenv

ROOT_PATH = Path(__file__).parent.parent

load_dotenv(ROOT_PATH / ".env")

DATA_PATH = ROOT_PATH / "data"

DB_PATH = DATA_PATH / "db.sqlite"
IMAGE_PATH = DATA_PATH / "images"


KEYCLOAK_URL = os.getenv("KEYCLOAK_URL")
KEYCLOAK_REALM = os.getenv("KEYCLOAK_REALM", "master")
