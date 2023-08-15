FROM python:3.8-slim

LABEL maintainer="Rafael Ferreira <rafael_low@hotmail.com>"
LABEL description "Dockerfile para criar imagem do container"

WORKDIR /app
COPY . /app
RUN pip3 install -r requirements.txt

CMD python3 exporter.py
