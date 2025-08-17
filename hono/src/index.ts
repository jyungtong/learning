import { Hono, MiddlewareHandler } from 'hono';
import user from './user';

const app = new Hono();


const logger: MiddlewareHandler = async (c, next) => {
  const startTime = new Date().getTime();
  console.log(`[${c.req.method}] ${c.req.url}`);
  await next();

  const endTime = new Date().getTime();
  console.log(`duration: ${endTime - startTime}`);
};

app.use('*', logger);
app.get('/', c => c.text(`hello from hono: ${new Date()}`));
app.get('/json', c => c.json({ msg: 'hello from hono json api' }));

app.route('/users', user);

export default app;
