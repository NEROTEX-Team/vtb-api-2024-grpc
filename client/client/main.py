from client.presentors.rest.config import RestConfig
from client.presentors.rest.factory import RestService

config = RestConfig()
app = RestService(config=config).create_application()
