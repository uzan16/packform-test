export interface PaginationRequestDto {
  pageSize: number;
  pageNumber: number;
  sortBy?: string;
  sortDirection?: string;
}
