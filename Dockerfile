FROM python:3.8-alpine as api1
COPY src/api1.py .
RUN pip install flask
CMD ["python", "api1.py"]

FROM python:3.8-alpine as api2
COPY src/api2.py .
RUN pip install flask
CMD ["python", "api2.py"]
