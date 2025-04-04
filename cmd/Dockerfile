FROM quay.io/psschwei/base-agent:latest

WORKDIR /usr/src/app

COPY requirements.txt .
RUN pip install -r requirements.txt

COPY . ./

CMD ["fastapi", "run", "agent.py", "--port", "8080"]
