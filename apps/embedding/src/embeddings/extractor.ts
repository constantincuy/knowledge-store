import { pipeline, FeatureExtractionPipeline } from '@xenova/transformers';

export class Extractor {
    public readonly model: string;
    private extractor: FeatureExtractionPipeline | null = null;
    private quantized: boolean = false;

    constructor(model: string) {
        this.model = model;
    }

    public useQuantized(val: boolean) {
        this.quantized = val;
    }

    public async preload() {
        if (!this.extractor) {
            this.extractor = await pipeline('feature-extraction', this.model, {
                quantized: this.quantized
            });
        }
    }

    public async get(): Promise<FeatureExtractionPipeline> {
        await this.preload();

        if (!this.extractor) {
            throw new Error("Model " + this.model + " could not be preloaded!")
        }

        return this.extractor;
    }

}