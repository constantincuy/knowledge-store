"use client"

import {Search} from "@/app/components/search";
import {Button} from "@/components/ui/button";
import {useCallback, useState} from "react";
import {useFetch} from "@/core/hooks/useFetch";
import {Card, CardContent, CardHeader} from "@/components/ui/card";
import {File} from "@/core/api/model/file";

export const SearchArea = () => {
    const [searching, setSearching] = useState(false)
    const [query, setQuery] = useState("")

    const {loading, error, data} = useFetch<File[]>({
        path: `http://localhost:8765/knowledge-base/development/files?q=${encodeURIComponent(query)}`,
        enabled: searching,
    });

    const searchDocs = useCallback(() => {
        setSearching(true)
    }, [setSearching])

    return (
        <div className="flex-1 space-y-8 p-8 pt-6">
            <div className="flex flex-col items-center justify-center mt-12 space-y-8">
                <h2 className="text-3xl font-bold tracking-tight">Search your documents</h2>
                <div className="w-1/2 flex gap-4">
                    <Search value={query} onChange={(e) => setQuery(e.target.value)}/>
                    <Button onClick={searchDocs}>
                        Search
                    </Button>
                </div>
            </div>
            <div className="space-y-4">
                {
                    !loading && !error && data && data.map((f, i) => (
                        <Card key={f.id + i}>
                            <CardHeader>
                                {f.path}
                            </CardHeader>
                            <CardContent>
                                <p className="text-sm text-muted-foreground">{f.provider}</p>
                            </CardContent>
                        </Card>
                    ))
                }
            </div>
        </div>
    )
}