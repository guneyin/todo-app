import { expect, test } from '@playwright/test';
import { provider } from './setup';

test.describe('Add item', () => {
	test.setTimeout(10000);

	test.beforeAll(async () => {
		await provider.setup();
	});

	test.afterAll(async () => {
		await provider.finalize().then(() => {
			provider.writePact();
		});
	});

	test('single item expecting to added to list', async ({ page }) => {
		await page.goto('/');

		const itemInput = page.getByTestId('add-item-input');
		await expect(itemInput).toBeVisible();
		await itemInput.fill('buy some milk');

		await provider.addInteraction({
			state: 'empty todo list',
			uponReceiving: 'add todo item',
			withRequest: {
				path: '/api/todos',
				method: 'POST',
				body: JSON.stringify({ todo: 'buy some milk' })
			},
			willRespondWith: {
				status: 200
			}
		});

		const addItemBtn = page.getByTestId('add-item-btn');
		await addItemBtn.click();

		await page.waitForTimeout(1000);

		const items = await page.getByTestId('todos').locator('ul > li').all();
		expect(items.length).toBe(2);

		const item = page.getByTestId('todos').locator('ul > li').last();
		await expect(item).toHaveText('buy some milk');

		await page.close();
	});
});
