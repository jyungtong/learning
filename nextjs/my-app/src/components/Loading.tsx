'use client'

export default function Loading({ children }: { children: React.ReactNode }) {
    return (
        <>
            <h1>...Loading</h1>

            {children}
        </>
    );
}
