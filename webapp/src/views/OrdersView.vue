<template>
  <div class="d-flex flex-column ga-6">
    <general-search v-model="search" />
    <v-row>
      <v-col cols="4">
        <date-range-picker
          v-model:startDate="dateRangeStart"
          v-model:endDate="dateRangeEnd"
          label="Created Date"
        />
      </v-col>
    </v-row>
    <span class="text-body-1">Total Amount: </span>
    <v-data-table-server
      v-model:items-per-page="itemsPerPage"
      :headers="headers"
      :items-length="totalItems"
      :items="data"
      :loading="loading"
      :page="currentPage"
      color="purple"
      @update:options="loadItems"
    >
      <template v-slot:[`item.orderName`]="{ value, item }">
        <div class="d-flex flex-column my-2">
          <span class="font-weight-bold">{{ value }}</span>
          <span>{{ item.products }}</span>
        </div>
      </template>
      <template v-slot:[`item.createdAt`]="{ value }">
        <span>{{ formatDate(value) }}</span>
      </template>
      <template v-slot:[`item.deliveredAmount`]="{ value }">
        <span>{{ formatCurrency(value) }}</span>
      </template>
      <template v-slot:[`item.totalAmount`]="{ value }">
        <span>{{ formatCurrency(value) }}</span>
      </template>

      <template v-slot:bottom="{ page, pageCount, itemsPerPage, setItemsPerPage }">
        <v-divider />
        <custom-pagination
          :page-count="pageCount"
          :page="page"
          :items-length="totalItems"
          :items-per-page="itemsPerPage"
          @update:items-per-page="setItemsPerPage"
          @update:page="(e) => (currentPage = e)"
        />
      </template>
    </v-data-table-server>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import generalSearch from '@/components/GeneralSearch.vue'
import dateRangePicker from '@/components/DateRangePicker.vue'
import customPagination from '@/components/CustomPagination.vue'
import type { VDataTable } from 'vuetify/components'
import { getOrders } from '@/service/ordersService'
import { formatDate, formatCurrency } from '@/helpers/format'

type ReadonlyHeaders = VDataTable['headers']

const headers: ReadonlyHeaders = [
  { title: 'Order Name', key: 'orderName' },
  { title: 'Customer Company', key: 'companyName' },
  { title: 'Customer Name', key: 'customerName' },
  { title: 'Order Date', key: 'createdAt' },
  { title: 'Delivered Amount', key: 'deliveredAmount', align: 'end' },
  { title: 'Total Amount', key: 'totalAmount', align: 'end' }
]
const itemsPerPage = ref<number>(5)
const totalItems = ref<number>(0)
const data = ref<any[]>([])
const loading = ref<boolean>(false)
const search = ref<string>('')
const dateRangeStart = ref<Date | undefined>()
const dateRangeEnd = ref<Date | undefined>()
const tableParam = ref()
const currentPage = ref(1)
let debounce: any

const loadItems = (params: any) => {
  tableParam.value = params
  getData()
}

const getData = async () => {
  try {
    loading.value = true
    const { itemsPerPage: pageSize, page: pageNumber, sortBy: sortObj } = tableParam.value
    let sortBy: string | undefined, sortDirection: string | undefined
    if (sortObj.length > 0) {
      sortBy = sortObj[0].key
      sortDirection = sortObj[0].order
    }
    const {
      success,
      data: remoteData,
      totalRows
    } = await getOrders({
      pageNumber,
      pageSize,
      sortBy,
      sortDirection,
      q: search.value,
      startDate: dateRangeStart.value,
      endDate: dateRangeEnd.value
    })
    if (!success) return
    totalItems.value = totalRows
    data.value = remoteData
  } finally {
    loading.value = false
  }
}

const searchChange = () => {
  if (debounce) clearTimeout(debounce)

  debounce = setTimeout(() => {
    getData()
  }, 750)
}

watch(search, searchChange)
watch([dateRangeStart, dateRangeEnd], (newVal) => {
  if (!newVal[0] === !newVal[1]) {
    // XNOR operator
    getData()
  }
})
</script>
