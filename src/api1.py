import flask
from flask import jsonify

app = flask.Flask(__name__)

@app.route('/')
def home():
    return jsonify({
        "message": "Hello from API #1!"
    })

if __name__ == '__main__':
    host = '0.0.0.0'
    port = 5000
    app.run(host=host, port=port, debug=True)