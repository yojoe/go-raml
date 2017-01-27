from sanic import Sanic
from sanic.response import json

from helloworld_if import helloworld_if


app = Sanic(__name__)

app.blueprint(helloworld_if)


app.static('/apidocs', './apidocs')
app.static('/', './index.html')

if __name__ == "__main__":
    app.run(debug=True, port=5000, workers=2)
