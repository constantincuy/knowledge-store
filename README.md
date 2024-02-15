# Knowledge store

Proof of concept for an AI based document indexing system using Sentence-Transformer written in Go and Typescript.

## Services
- embedding service - generates text embeddings [embedding](./apps/embedding)
- store service - Collects Document files from different file storage provider and indexes them in a Vector database [store](./apps/store)

## Setup
Run `npm i` in the root of the project to install all required dependencies. After installing the dependencies run
`npm run docker:build` to build the container images of the store service and embedding service. Next copy the [example-compose.yml](example-compose.yml)
to a new `compose.yml` and run `docker compose up (-d)`.

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

