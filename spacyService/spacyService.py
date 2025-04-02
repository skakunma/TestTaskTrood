from flask import Flask, request, jsonify
import spacy
import os
from dotenv import load_dotenv

load_dotenv()

HOST_SERVER = os.getenv("HOST_SPACY")
PORT_SERVER = os.getenv("PORT_SPACY")


nlp = spacy.load("en_core_web_sm")

app = Flask(__name__)

@app.route('/process', methods=['POST'])
def process_text():
    data = request.get_json()
    text = data.get("text", "")
    
    doc = nlp(text)
    
    entities = [{"text": ent.text, "label": ent.label_} for ent in doc.ents]
    
    return jsonify({"entities": entities})


if __name__ == '__main__':
    app.run(debug=True, host=HOST_SERVER, port=PORT_SERVER)