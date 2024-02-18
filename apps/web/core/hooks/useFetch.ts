import {useEffect, useState} from "react";

export interface FetchResult<Type> {
    loading: boolean
    data?: Type
    error?: Error
    abort: () => void
}

export interface FetchOptions {
    enabled: boolean
    path: string
    req?: RequestInit
}

export const useFetch = <Result>({path, req, enabled}: FetchOptions): FetchResult<Result> => {
    const [controller, setController] = useState<AbortController | undefined>(undefined)
    const [loading, setLoading] = useState(true)
    const [data, setData] = useState(undefined)
    const [error, setError] = useState<Error | undefined>(undefined)

    useEffect(() => {
        setController(new AbortController())
    }, [setController]);

    useEffect(() => {
        console.log(enabled && controller)
        if (enabled && controller) {
            const signal = controller.signal;
            fetch(path, req ? {...req, signal} : {signal}).then(res => {
                setLoading(false);
                if (!res.ok) {
                    setError(new Error(`${res.status}: ${res.statusText}`));
                    return;
                }

                res.json().then(data => {
                    setData(data)
                })
            }).catch(err => {
                setLoading(false)
                setError(err)
            })
        }

    }, [path, req, enabled, controller, setLoading, setData, setError]);

    return {
        loading,
        data,
        error,
        abort: () => controller?.abort()
    }
}