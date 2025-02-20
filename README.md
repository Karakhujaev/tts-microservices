Text-to-Speech Microservices

Overview

Implementation of a Text-to-Speech (TTS) microservice-based system that:

Receives text input via a REST API.

Stores text in a PostgreSQL database.

Sends text to a Speech Service for conversion.

Uses Coqui TTS to generate speech.

Stores the generated audio files and provides an API to retrieve them.

Deploys using Kubernetes and is accessible via NGINX API Gateway.


Architecture

```bash
📁 tts-microservices/
│── 📁 text-service/        
│── 📁 speech-service/      
│── 📁 k8s-configs/         
│── docker-compose.yml      
│── README.md
```

Technologies Used

Go – Backend for both services.

Gin – Web framework for RESTful APIs.

PostgreSQL – Database for storing text.

Coqui TTS – Converts text to speech.

Docker – Containerized services.

Kubernetes – Deployment and scaling.

NGINX – API Gateway for routing.

Installation & Setup

1) Clone the Repository

```bash
git clone https://github.com/Karakhujaev/tts-microservices.git
cd tts-microservices
```

2) Set Up Environment Variables

Create .env files in text-service/ and speech-service/.

text-service/.env

```bash
DATABASE_URL=postgres://user:password@db:5432/textdb
TEXT_SERVICE_PORT=8080
SPEECH_SERVICE_URL=http://speech-service:8081
```

speech-service/.env
```bash
SPEECH_SERVICE_PORT=8081
AUDIO_STORAGE_PATH=./audio
```

Running Docker Compose
```bash
docker-compose up --build
```


Deployment to Kubernetes
```bash
kubectl apply -f k8s-configs/namespace.yaml
kubectl apply -f k8s-configs/secrets.yaml
kubectl apply -f k8s-configs/volumes.yaml
kubectl apply -f k8s-configs/postgres.yaml
kubectl apply -f k8s-configs/text-service.yaml
kubectl apply -f k8s-configs/speech-service.yaml
kubectl apply -f k8s-configs/nginx.yaml
```

Verify Services
```bash
kubectl get pods -n tts
kubectl get services -n tts
```

Access the API

Find the external IP of the NGINX API Gateway:
```bash
kubectl get svc nginx -n tts
```


API Documentation

1) Submit Text for Processing

Endpoint: POST /api/text
```bash
{
  "text": "Hello, this is a test message"
}
```

Response:
```bash
{
  "id": "123",
  "text": "Hello, this is a test message",
  "status": "processing"
}
```

2) Retrieve Processed Speech

Endpoint: GET /api/speech/{id}

Response:
```bash
{
  "id": "123",
  "text": "Hello, this is a test message",
  "audio_url": "/audio/123.mp3",
  "status": "completed"
}
```



```mermaid
graph TD;
    Client -->|API Request| NGINX["NGINX API Gateway (80)"]
    NGINX -->|POST /api/text| TextService["Text Service (8080)"]
    NGINX -->|POST /process| SpeechService["Speech Service (8081)"]
    TextService -->|Stores text| PostgreSQL["PostgreSQL DB (5432)"]
    TextService -->|Forwards text| SpeechService
    SpeechService -->|Converts text to audio| Storage["Audio Storage"]
    NGINX -->|GET /api/speech/{id}| Storage
```