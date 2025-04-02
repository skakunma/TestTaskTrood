FROM python:3.9-slim

WORKDIR /app

COPY . .

RUN apt-get update && apt-get install -y \
    gcc \
    g++ \
    python3-dev \
    build-essential \
    libatlas-base-dev

RUN pip install --upgrade pip 

RUN pip install -U pip setuptools wheel

RUN pip install -r requirements.txt


RUN python -m spacy download en_core_web_sm

EXPOSE 5000

CMD ["python", "spacyService.py"]