import type { OrdersRequestDto, OrdersResponseDto } from '@/dto/ordersDto'
import type { SuccessPaginateResponse } from '.'
import { formatObjectToQuery } from '@/helpers/format'

const BASE_API_URL = '/api'

export const getOrders = async (
  param: OrdersRequestDto
): Promise<SuccessPaginateResponse<OrdersResponseDto>> => {
  try {
    const response = await fetch(`${BASE_API_URL}/orders?${formatObjectToQuery(param)}`)
    return await response.json()
  } catch (error) {
    // handle the error
    console.error('error', error)
    throw error
  }
}
