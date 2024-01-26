export interface SuccessResponse {
  success: boolean
  message: string
}

export interface SuccessPaginateResponse<T> extends SuccessResponse {
  data: Array<T>
  totalRows: number
}
