# Knowledge store

Proof of concept for an AI based document indexing system using Sentence-Transformer written in Go and Typescript.

## Services
- embedding service - generates text embeddings [embedding](./apps/embedding)
- store service - Collects Document files from different file storage provider and indexes them in a Vector database [store](./apps/store)

## Setup
First setup a docker container with Postgres and PgVector:
```shell
docker run -p 5432:5432 -e POSTGRES_PASSWORD=root -e POSTGRES_USER=root -d ankane/pgvector 
```
(DB user and pw are hardcoded. CLI flags and env options will be added later)

Start the embedding server by either running `npm run dev` in `/apps/embedding` or in the root folder of this monorepo.
Hint: The embedding server downloads the `nomic-ai/nomic-embed-text-v1` [HuggingFace](https://huggingface.co/nomic-ai/nomic-embed-text-v1) automatically this can take a few minutes on the first run.

Lastly build and run the store server by running
```shell
cd apps/store
go run .\cmd\web\main.go
```

Creating a knowledge base:
```
curl --location 'http://localhost:8765/knowledge-base/' \
--header 'Content-Type: application/json' \
--data '{
    "name": "development"
}'
```

Creating a new knowledge base will spawn new worker threads which will listen for file changes
on the configured storage providers. For demonstration purposes there is only a fake storage at the moment
which will index 3 hard coded files: hund.txt, katze.txt and flugzeug.txt (content was taken from Wikipedia).
(Indexing file check currently runs every 10 seconds and watches for file changes on the storage provider)

Querying the knowledge base:
```
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

