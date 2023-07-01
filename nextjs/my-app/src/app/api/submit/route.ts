import { NextResponse } from 'next/server';

export function POST(req: Request) {
    // Get data submitted in request's body.
    const body = req.body;

    // Optional logging to see the responses
    // in the command line where next.js app is running.
    console.log('body: ', body);

    // Guard clause checks for first and last name,
    // and returns early if they are not found
    if (!body || !body.first || !body.last) {
        // Sends a HTTP bad request error code
        return NextResponse.json({ data: 'First or last name not found' }, { status: 400 });
    }

    // Found the name.
    // Sends a HTTP success code
    return NextResponse.json({ data: `${body.first} ${body.last}` }, { status: 200 });
}
