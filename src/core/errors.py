class DomainError(Exception):
    """Base class for domain errors."""

    def __init__(self, message: str, http_status_code: int = 400):
        super().__init__(message)

        self.http_status_code = http_status_code
