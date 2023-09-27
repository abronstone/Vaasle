FROM python:3.8-alpine as engine
COPY src/engine/* .
RUN pip install flask
CMD ["python", "engine.py"]

FROM python:3.8-alpine as play-game
COPY src/play-game/* .
RUN pip install flask
CMD ["python", "play-game.py"]
