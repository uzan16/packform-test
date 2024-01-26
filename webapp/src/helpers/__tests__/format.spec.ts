import { describe, test, expect } from 'vitest'

import {
  formatCurrency,
  formatDate,
  formatObjectToQuery
} from '../format'

describe('helper format', () => {
  test('formatCurrency', () => {
    const param = 1234.54;
    const res = formatCurrency(param)
    expect(res).toBe("$ 1,234.54")
  })

  test('formatDate', () => {
    const expectTime1 = "Jan 1st 2024, 12:00 AM";
    const expectTime2 = "Jan 1st 2024, 4:00 AM";

    const resMelbourneTime = formatDate(new Date('2024-01-01T00:00:00.000+11:00'))
    expect(resMelbourneTime).toBe(expectTime1)

    const resJakartaTime = formatDate(new Date('2024-01-01T00:00:00.000+07:00'))
    expect(resJakartaTime).toBe(expectTime2)
  })

  test('formatObjectToQuery', () => {
    const expectedVal = "one=val1&two=val2&three=3&four=2023-12-31T13:00:00.000Z&";

    const regex = new Date().toString().match(/([-+][0-9]+)\s/);
    const offset = regex ? regex[1] : "+0000";
    const res = formatObjectToQuery({
      one: "val1",
      two: "val2",
      three: 3,
      four: new Date(`2024-01-01T00:00:00.000${offset}`)
    })

    expect(res).toBe(expectedVal)

  })
})
