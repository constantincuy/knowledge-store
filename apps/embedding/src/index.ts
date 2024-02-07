import {Extractor} from "./embeddings/extractor.js";
import Fastify from 'fastify'
import {EmbeddingBuilder} from "./embeddings/embedding-builder.js";
const fastify = Fastify({
    logger: true
});

fastify.log.info("Preloading AI model...")
const extractor = new Extractor('nomic-ai/nomic-embed-text-v1');

extractor.preload().then(() => {
    fastify.log.info(`Loaded model ${extractor.model}`)

    fastify.get('/', async (request, reply) => {
        reply.type('application/json').code(200)
        return { service: 'embedding-server', loadedModel: extractor.model }
    })


    fastify.post('/embeddings/generate', async (request, reply) => {
        if (request.headers["content-type"] !== 'application/json') {
            reply.type('application/json').code(400)
            return { error: 'Invalid content-type header only "application/json" is supported!' }
        }


        const texts: any = request.body;
        if (!Array.isArray(texts)) {
            reply.type('application/json').code(400)
            return { error: 'Invalid content in request provide array of strings!' }
        }

        const builder = new EmbeddingBuilder();
        texts.forEach(t => builder.addText(t))

        reply.type('application/json').code(200)
        return builder.buildWith(extractor);
    })

    fastify.listen({ port: 3000 }, (err) => {
        if (err) throw err
    })
}).catch(e => {
    fastify.log.error(`Could not start server: ${e}`)
})

