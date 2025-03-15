import { expect, test } from '@playwright/test';

test.describe('Consumer tests', () => {
	test.setTimeout(10000);

	test.describe('Add item', async () => {
		test('two items expecting to added to list', async ({ page }) => {
			await page.goto('/');

			const itemInput = page.getByTestId('itemInput');
			await expect(itemInput).toBeVisible()
			await itemInput.fill('buy some milk');

			const addItemBtn = page.getByTestId('addItemBtn');
			await addItemBtn.click();

			await page.waitForTimeout(1000);

			await itemInput.fill('buy fruits');
			await addItemBtn.click();

			await page.waitForTimeout(1000);

			const items = await page.getByTestId('todos').locator('li').all();
			expect(items.length).toBe(2);

			const item = page.getByTestId('todos').locator('ul > li').first();
			await expect(item).toHaveText('buy some milk');
		});
	})
});
