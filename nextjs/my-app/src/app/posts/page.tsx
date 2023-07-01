import Link from 'next/link';
import Loading from '@src/components/Loading';

type Post = {
    id: string;
    userId: string;
    title: string;
    body: string;
};

async function getPosts() {
    const response = await fetch(
        `https://jsonplaceholder.typicode.com/posts?API_KEY=${process.env.API_KEY}`
    );

    return response.json();
}

export default async function Page() {
    const data = await getPosts();

    return (
        <>
            <Loading>
                {data.map((d: Post) => {
                    return (
                        <Link
                            key={d.id}
                            href={`/posts/${d.id}`}
                            prefetch={false}
                        >
                            <h1>{d.title}</h1>
                        </Link>
                    );
                })}
            </Loading>
        </>
    );
}
