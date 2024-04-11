import { it, expect } from 'vitest';
import supertest from 'supertest';
import { Api } from 'sst/node/api';

const request = supertest(Api.api.url);

it('should be true', async () => {
    const res = await request.get('/');

    expect(res.status).toEqual(200);
    expect(res.text).toEqual(expect.stringContaining('Hello world'));
});
