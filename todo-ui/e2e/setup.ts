import { Pact } from '@pact-foundation/pact';
import path from 'node:path';

export const provider = new Pact({
	port: 3001,
	dir: path.resolve(process.cwd(), 'pacts'),
	consumer: 'todo-ui',
	provider: 'todo-api',
	pactfileWriteMode: 'merge'
});
