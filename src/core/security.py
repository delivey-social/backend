import jwt
from fastapi import Depends, HTTPException
from fastapi.security import HTTPAuthorizationCredentials, HTTPBearer
from jwt import PyJWKClient  # type: ignore[attr-defined]
from pydantic import BaseModel

from src.constants import KEYCLOAK_REALM, KEYCLOAK_URL

JWKS_URL = f"{KEYCLOAK_URL}/realms/{KEYCLOAK_REALM}/protocol/openid-connect/certs"
EXPECTED_ISSUER = f"{KEYCLOAK_URL}/realms/{KEYCLOAK_REALM}"

jwks_client = PyJWKClient(JWKS_URL)

security = HTTPBearer()


class CurrentUser(BaseModel):
    id: str
    email: str | None = None
    username: str | None = None
    roles: list[str] = []


def get_current_user(
    credentials: HTTPAuthorizationCredentials = Depends(security),
) -> CurrentUser:
    print("[DEBUG] get_current_user called with credentials:", credentials)

    token = credentials.credentials

    try:
        signing_key = jwks_client.get_signing_key_from_jwt(token).key

        payload = jwt.decode(
            token,
            signing_key,
            algorithms=["RS256"],
            issuer=EXPECTED_ISSUER,
            options={"verify_aud": False},  # Disable audience verification
        )

        print("[DEBUG] JWT payload:", payload)

    except jwt.ExpiredSignatureError:
        raise HTTPException(
            status_code=401,
            detail="Token has expired",
        )
    except jwt.InvalidTokenError as e:
        print("[DEBUG] Invalid token error occurred")
        print("[DEBUG] Error details:", e)

        raise HTTPException(
            status_code=401,
            detail="Invalid token",
        )
    except jwt.PyJWKClientError as e:
        print("[DEBUG] PyJWKClientError occurred")
        print("[DEBUG] Error details:", e)

        raise HTTPException(
            status_code=503,
            detail="Error connecting to JWKS endpoint",
        )

    user_id = payload.get("sub")

    if not user_id:
        raise HTTPException(
            status_code=401,
            detail="Token does not contain a valid user ID",
        )

    return CurrentUser(
        id=user_id,
        email=payload.get("email"),
        username=payload.get("preferred_username"),
        roles=payload.get("realm_access", {}).get("roles", []),
    )


def admin_protected_route(user: CurrentUser = Depends(get_current_user)):
    if "admin" not in user.roles:
        raise HTTPException(
            status_code=403,
            detail="You do not have permission to access this resource",
        )
