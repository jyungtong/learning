async function getPostById(postId: string) {
    const response = await fetch(
        `https://jsonplaceholder.typicode.com/posts/${postId}`,
        {
            cache: 'no-store',
        }
    );

    return response.json();
}

export default async function Page({
    params,
}: {
    params: { post_id: string };
}) {
    const data = await getPostById(params.post_id);

    return (
        <>
            <h1>{data.title}</h1>
            <h2>{data.body}</h2>
        </>
    );
}
