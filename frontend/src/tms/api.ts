const DEFAULT_BASE_URL = "http://localhost:8000/api/v1"

export async function api<T>(path: string, opts?: RequestInit & { baseUrl?: string }) {
    const baseUrl = opts?.baseUrl ?? DEFAULT_BASE_URL
    const res = await fetch(`${baseUrl}${path}`, {
        headers: { "Content-Type": "application/json" },
        ...opts
    })
    if (!res.ok) {
        const text = await res.text()
        throw new Error(`${res.status} ${res.statusText}: ${text}`)
    }
    return (await res.json()) as T
}

export { DEFAULT_BASE_URL }
