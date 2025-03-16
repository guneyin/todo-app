import type { Actions, PageServerLoad } from './$types';
import { env } from '$env/dynamic/public';

export const load: PageServerLoad = async () => {
	let todos: string[] = [];

	const fetched = await fetch(`http://${env.PUBLIC_API_URL}/api/todos`)
		.then((res) => res.json())
		.catch((err) => console.log(err));

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
		await fetch(`http://${env.PUBLIC_API_URL}/api/todos`, {
			method: 'POST',
			body: body,
			headers: {
				'Content-Type': 'application/json'
			}
		}).catch((err) => {
			throw err;
		});
	}
};
