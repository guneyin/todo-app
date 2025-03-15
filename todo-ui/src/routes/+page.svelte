<script lang="ts">
	import { enhance } from '$app/forms';

	let todo = $state('');
	let todos = $state([]) as string[];

	let { data } = $props();

	data.todos.forEach((el: string) => {
		todos.push(el);
	});
</script>

<article data-testid="add-item">
	<form
		method="POST"
		action="?/addItem"
		use:enhance={() => {
			return async ({ result }) => {
				if (result.type === 'success') {
					todos.push(todo.trim());
				}
			};
		}}
	>
	<fieldset role="group">
		<input
			data-testid="add-item-input"
			id="todo-item"
			name="todo-item"
			type="text"
			placeholder="Add new todo"
			aria-label="Add new todo"
			bind:value={todo}
		/>
		<input data-testid="add-item-btn" type="submit" value="+Add" />
	</fieldset>
	</form>
</article>

<article data-testid="todos">
	{#if todos.length === 0}
		<span>No data found</span>
	{:else}
		<ul>
			{#each todos as item}
				<li>{item}</li>
			{/each}
		</ul>
	{/if}
</article>
