import { test, expect, Locator } from '@playwright/test';

const waitForLoading = async(p: Locator) => {
  await p.locator(".v-data-table--loading").waitFor({
    state: 'visible'
  });
  await p.locator(".v-data-table--loading").waitFor({
    state: 'detached'
  });
}

test.describe("Test Orders Page", () => {
  test("'/' route to '/orders'", async ({ page }) => {
    await page.goto('/');
    await page.getByTestId('view-orders').first().waitFor({
      state: 'visible'
    });
    await expect(page.url()).toContain("/orders")
  })
  test("'/anything' route to '/orders'", async ({ page }) => {
    await page.goto('/anything');
    await page.getByTestId('view-orders').first().waitFor({
      state: 'visible'
    });
    await expect(page.url()).toContain("/orders")
  })
  test("test functionality", async ({ page }) => {
    await page.goto('/orders');
    await page.getByTestId('view-orders').first().waitFor({
      state: 'visible'
    });
    
    const currentPage = page.getByTestId('view-orders');
    const searchBar = currentPage.getByTestId('general-search');
    const dateRangePicker = currentPage.getByTestId('date-range-picker');
    const startDatePicker = dateRangePicker.getByTestId('date-range-picker-start-date');
    const endDatePicker = dateRangePicker.getByTestId('date-range-picker-end-date');
    const calendarStart = page.getByTestId('date-range-picker-calendar-start');
    const calendarEnd = page.getByTestId('date-range-picker-calendar-end');
    const table = currentPage.getByTestId('data-table');
    const pagination = currentPage.getByTestId('pagination');
    const pageNumber = pagination.getByTestId('input-page-number').locator("input");
    
    await pagination.getByLabel('Next page').click();
    await expect(pageNumber).toHaveValue("2");
    await waitForLoading(currentPage);

    await pagination.getByLabel('Previous page').click();
    await expect(pageNumber).toHaveValue("1");
    await waitForLoading(currentPage);

    const itemPerPage = pagination.getByTestId('select-item-per-page');
    await itemPerPage.click();
    await page.getByText('10 / page').click();
    await expect(itemPerPage).toHaveText("10 / page");
    await waitForLoading(currentPage);

    await itemPerPage.click();
    await page.getByText('5 / page').click();
    await expect(itemPerPage).toHaveText("5 / page");
    await waitForLoading(currentPage);
    
    let count = await table.locator('.v-data-table__tr').count();
    await expect(count).toBeLessThanOrEqual(5);

    await searchBar.locator('input').fill('just a random text blablabla');
    await waitForLoading(currentPage);
    await expect(table.locator(".v-data-table-rows-no-data")).toBeVisible();
    
    await searchBar.getByRole('button', { name: 'Clear' }).click();
    await waitForLoading(currentPage);
    count = await table.locator('.v-data-table__tr').count();
    await expect(count).toBeLessThanOrEqual(5);
    
    await startDatePicker.locator('input').click();
    await calendarStart.getByRole('button', { name: '1', exact: true }).click();
    await endDatePicker.locator('input').click();
    await calendarEnd.getByRole('button', { name: '2', exact: true }).click();
    await waitForLoading(currentPage);
    await expect(table.locator(".v-data-table-rows-no-data")).toBeVisible();

    await startDatePicker.locator('input').click();
    await calendarStart.locator('div').filter({ hasText: /^January 2024$/ }).getByRole('button').nth(1).click();
    await calendarStart.getByRole('button', { name: '2010' }).click();
    await calendarStart.getByRole('button', { name: '1', exact: true }).click();
    await waitForLoading(currentPage);

    count = await table.locator('.v-data-table__tr').count();
    await expect(count).toBeLessThanOrEqual(5);
  })
});
