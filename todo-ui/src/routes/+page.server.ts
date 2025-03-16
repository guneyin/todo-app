import type { Actions, PageServerLoad } from './$types';
import { env } from '$env/dynamic/public';
import { fail } from '@sveltejs/kit';

export const load: PageServerLoad = async () => {
	let todos: string[] = [];

	const fetched = await fetch(`http://${env.PUBLIC_API_URL}/api/todos`)
		.then(async (r) => {
			if (!r.ok) {
				const body = await r.text();
				return fail(r.status, { error: body });
			}

			return await r.json();
		})
		.catch((err) => fail(500, { error: err.error }));

	if (fetched) {
		todos = fetched.todos;
	}

	return { todos };
};

export const actions: Actions = {
	addItem: async ({ request }) => {
		const data = await request.formData();
		const todo = data.get('todo-item');

		const body = JSON.stringify({ todo: todo });
		return await fetch(`http://${env.PUBLIC_API_URL}/api/todos`, {
			method: 'POST',
			body: body,
			headers: {
				'Content-Type': 'application/json'
			}
		})
			.then(async (r) => {
				if (!r.ok) {
					const body = await r.text();
					return fail(r.status, { error: body });
				}
			})
			.catch((err) => fail(500, { error: err.error }));
	}
};
