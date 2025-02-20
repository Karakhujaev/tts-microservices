FROM python:3.9

RUN apt update && apt install -y ffmpeg && pip install tts

WORKDIR /app
COPY . .

RUN pip install -r requirements.txt

CMD ["python", "main.py"]

EXPOSE 8081
