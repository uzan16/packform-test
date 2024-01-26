import dayjs from 'dayjs'
import utc from 'dayjs/plugin/utc'
import timezone from 'dayjs/plugin/timezone'
dayjs.extend(utc)
dayjs.extend(timezone)

const nthNumber = (num: number) => {
  if (num > 3 && num < 21) return 'th'
  switch (num % 10) {
    case 1:
      return 'st'
    case 2:
      return 'nd'
    case 3:
      return 'rd'
    default:
      return 'th'
  }
}

export const formatCurrency = (num: number) => {
  return `$ ${num.toFixed(2).replace(/\d(?=(\d{3})+\.)/g, '$&,')}`
}

export const formatDate = (date: Date) => {
  const djs = dayjs(date).tz('Australia/Melbourne')
  const day = djs.format('D')
  return `${djs.format('MMM D')}${nthNumber(+day)} ${djs.format('YYYY, h:mm A')}`
}

export const formatObjectToQuery = <T>(obj: T) => {
  if (typeof obj !== 'object') return ''

  let str = ''
  for (const key in obj) {
    let val: any = obj[key]
    if (val instanceof Date) {
      val = dayjs(val).tz('Australia/Melbourne', true).toISOString()
    }
    str += `${key}=${val ?? ''}&`
  }
  return str
}
