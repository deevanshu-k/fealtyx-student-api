# FealtyX Assignment

## Setup
- Create env file `.env` in root directory of the project
  ```env
    PORT = 8050
    OLLAMA_GENERATE_URL = "http://localhost:11434/api/generate"
    OLLAMA_MODEL = "llama3.2:1b"
  ```

## APIs
- Get All Students -> GET /students
- Get Student -> GET /student/:id
- Create Student -> POST /students
  ```json
  {
    "name": "string",
    "age": "number",
    "email": "string"
  }
  ```
- Update Student -> PUT /students/:id
  ```json
  {
    "name": "string", // OPTIONAL
    "age": "number",  // OPTIONAL
    "email": "string" // OPTIONAL
  }
  ```
- Delete Student -> DELETE /students/:id
- Get Student Summary -> GET /student/:id/summary
