export const fetcher = async (
    resource: RequestInfo,
    init?: RequestInit,
    // eslint-disable-next-line react-hooks/exhaustive-deps
): Promise<any> => {
    const res = await fetch(resource, init);

    if (!res.ok) {
        const errRes = await res.json();
        
        const err = new Error(
            errRes.message ?? 'APIリクエスト中にエラーが発生しました',
        )
        throw err
    }

    return res.json();
}