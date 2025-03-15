import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async () => {
	let todos: string[] = [];

	const fetched = await fetch('http://127.0.0.1:3001/api/todos')
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
		await fetch('http://127.0.0.1:3001/api/todos', {
			method: 'POST',
			body: body
		}).catch((err) => {
			throw err;
		});
	}
};
