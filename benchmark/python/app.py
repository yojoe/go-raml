from flask import Flask, send_from_directory, send_file
import wtforms_json
from users import users_api


app = Flask(__name__)

import logging
log = logging.getLogger('werkzeug')
log.setLevel(logging.ERROR)

app.config["WTF_CSRF_ENABLED"] = False
wtforms_json.init()

app.register_blueprint(users_api)



@app.route('/apidocs/<path:path>')
def send_js(path):
    return send_from_directory('apidocs', path)


@app.route('/', methods=['GET'])
def home():
    #return send_file('index.html')
    return "Hello World!"

if __name__ == "__main__":
    app.run(debug=False)
