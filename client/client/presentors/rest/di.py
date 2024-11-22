from pathlib import Path

from dishka import Provider, Scope, provide
from fastapi.templating import Jinja2Templates

PROJECT_PATH = Path(__file__).parent.parent.parent


class RestProvider(Provider):
    @provide(scope=Scope.APP)
    def templates(self) -> Jinja2Templates:
        return Jinja2Templates(directory=PROJECT_PATH / "presentors/rest/templates")
