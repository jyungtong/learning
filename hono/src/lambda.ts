import { handle } from 'hono/aws-lambda';
import app from '.';

export const handler = handle(app)
