#! /usr/bin/env python3
from fastapi import FastAPI
from pydantic import BaseModel
import uvicorn

class ModelInput(BaseModel):
    prompt: str

app = FastAPI()

@app.post("/")
async def run_agent(modelInput: ModelInput):
    prompt = modelInput.prompt

    # User-created agent

    ##### INSERT USER CODE HERE #####"

    # end user created content

    return {"output": messages}

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000, log_level="warning")
