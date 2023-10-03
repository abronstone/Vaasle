FROM python:3.8-alpine as play-game
COPY src/play-game/* .
RUN pip install flask
CMD ["python", "play-game.py"]