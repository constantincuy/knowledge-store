import {Extractor} from "./extractor.js";

export interface Embedding {
    dims: number[]
    type: string
    vectors: number[]
}

export class EmbeddingBuilder {
    private texts: string[] = [];
    private model: string = "nomic-ai/nomic-embed-text-v1";
    private pooling: 'none'|'mean'|'cls' = "mean";
    private normalize: boolean = true;


    public addText(text: string) {
        this.texts.push(text);
    }

    public usePooling(val: 'none' | 'mean' | 'cls') {
        this.pooling = val;
    }

    public useNormalize(val: boolean) {
        this.normalize = val;
    }

    public async buildWith(extractor: Extractor): Promise<Embedding> {
        const extract = await extractor.get();

        const result = await extract(this.texts, { pooling: this.pooling, normalize: this.normalize });

        return {
            dims: result.dims,
            type: result.type,
            vectors: [...result.data]
        }
    }


}