# Knowledge store

Proof of concept for an AI based document indexing system using Sentence-Transformer written in Go and Typescript.

## Services
- [Embedding service](./apps/embedding) - generates text embeddings 
- [Store service](./apps/embedding) - Collects Document files from different file storage provider and indexes them in a Vector database (Postgres with [pgvector](https://github.com/pgvector/pgvector))

### Setting up the Project

1. Run the following command in the root of the project to install all required dependencies:

    ```bash
    npm i
    ```

2. After installing the dependencies, build the container images for the store service and embedding service using the following command:

    ```bash
    npm run build:docker
    ```

3. Copy the [example-compose.yml](example-compose.yml) file to a new `compose.yml`.

4. Run the following command to start the services:

    ```bash
    docker compose up (-d)
    ```

    **Note:**
    - Depending on your internet speed, you may need to wait for a few minutes for the embedding service to download the `nomic-ai/nomic-embed-text-v1` model from [HuggingFace](https://huggingface.co/nomic-ai/nomic-embed-text-v1).
    - Before continuing, ensure that the following log lines appear in the terminal by running the following command:
      ```bash
      docker compose logs embedding

      {"level":30,"time":1708029322690,"pid":25,"hostname":"e2fe08cd66ed","msg":"Loaded model nomic-ai/nomic-embed-text-v1"}
      {"level":30,"time":1708029322704,"pid":25,"hostname":"e2fe08cd66ed","msg":"Server listening at http://0.0.0.0:3000"}
      ```

5. Creating a knowledge base:
    ```bash
    curl --location 'http://localhost:8765/knowledge-base/' \
    --header 'Content-Type: application/json' \
    --data '{
        "name": "development"
    }'
    ```
    **Note:**
    - Creating a new knowledge base will spawn new worker threads which will listen for file changes on the configured storage providers.
    - For demonstration purposes there is only a fake storage at the moment which will index 3 hard coded files: hund.txt, katze.txt and flugzeug.txt (content was taken from Wikipedia).
    - Indexing file check currently runs every 10 seconds and watches for file changes on the storage provider

6. Querying the knowledge base:
    ```bash
    curl --location 'http://localhost:8765/knowledge-base/development/files?q=Was%20kann%20fliegen%3F'
    ```

    Example response:
    ```json
    {
        "Files": [
            {
                "Id": "54f1591f-c2cb-42c7-8230-9305875020a9",
                "Path": "my/path/flugzeug.txt",
                "Provider": "fake_provider_1"
            }
          ]
    }
    ```

