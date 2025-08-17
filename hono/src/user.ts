import { Hono } from 'hono';

const user = new Hono();

const sleep = (ms) => new Promise(res => setTimeout(res, ms));

user.get('/', (c) => c.json({ users: ['Alice', 'Bob'] }))
user.get('/:id', async (c) => {
  // await sleep(1000);
  const id = c.req.param('id')
  return c.json({ id, name: `User ${id}` })
})

export default user;
