import type { PaginationRequestDto } from "./paginationDto";

export interface OrdersResponseDto {
  id: number;
  orderName: string;
  createdAt: Date;
  companyName: string;
  customerName: string;
  totalAmount: number;
  deliveredAmount: number;
  products: string;
}

export interface OrdersRequestDto extends PaginationRequestDto {
  q?: string;
  startDate?: Date;
  endDate?: Date;
}