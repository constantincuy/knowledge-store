import {File} from "@/core/api/model/file";

export async function search(query: string): Promise<File[]> {
    const res = await fetch(`http://localhost:8765/knowledge-base/development/files?q=${encodeURIComponent(query)}`)

    if (!res.ok) {
        throw new Error('Failed to fetch data')
    }

    return res.json()
}