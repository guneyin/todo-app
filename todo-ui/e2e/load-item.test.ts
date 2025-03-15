import { expect, test } from '@playwright/test';
import { provider } from './setup';

test.describe('Load item', () => {
	test.setTimeout(10000);

	test.beforeAll(async () => {
		await provider.setup();
		await provider.removeInteractions();
	});

	test.afterAll(async () => {
		await provider.finalize().then(() => {
			provider.writePact();
		});
	});

	test('single item expected to exist in the list', async ({ page }) => {
		await provider.removeInteractions();

		await provider.addInteraction({
			state: 'empty todo list',
			uponReceiving: 'fetch one item',
			withRequest: {
				path: '/api/todos',
				method: 'GET'
			},
			willRespondWith: {
				status: 200,
				body: JSON.stringify({ todos: ['buy some milk'] })
			}
		});

		await page.goto('/');

		await page.waitForTimeout(1000);

		const items = await page.getByTestId('todos').locator('ul > li').all();
		expect(items.length).toBe(1);

		await page.close();
	});
});
