from flask import Flask, request, Response
import json

app = Flask(__name__)

@app.route("/", methods=["GET"])
def home():
    response_data = {"message": "🎉 Welcome to the Flask API"}
    return Response(json.dumps(response_data), status=200, mimetype="application/json")

@app.route("/submit", methods=["POST"])
def submit():
    if not request.is_json:
        error_data = {"error": "Request must be JSON"}
        return Response(json.dumps(error_data), status=400, mimetype="application/json")

    data = request.get_json()
    if "name" not in data:
        error_data = {"error": "Missing 'name' field"}
        return Response(json.dumps(error_data), status=400, mimetype="application/json")

    success_data = {"message": f"✅ Hello, {data['name']}! Your data was received."}
    return Response(json.dumps(success_data), status=201, mimetype="application/json")

if __name__ == "__main__":
    app.run()
