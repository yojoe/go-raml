from flask import Flask
import wtforms_json
from deliveries import deliveries_api
from drones import drones_api


app = Flask(__name__)

app.config["WTF_CSRF_ENABLED"] = False
wtforms_json.init()

app.register_blueprint(deliveries_api)
app.register_blueprint(drones_api)


if __name__ == "__main__":
    app.run(debug=True)
