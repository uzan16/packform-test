<template>
  <div class="d-flex flex-row justify-center align-center mt-4 ga-4">
    <span class="text-body-2">Total: {{ props.itemsLength }}</span>
    <div class="d-flex flex-column">
      <v-select
        density="compact"
        :items="pageOptions"
        variant="outlined"
        hide-details
        :model-value="props.itemsPerPage"
        @update:model-value="(e) => emit('update:itemsPerPage', e)"
      ></v-select>
    </div>
    <v-pagination
      :length="props.pageCount"
      :total-visible="3"
      :model-value="props.page"
      @update:model-value="changePageNumber"
    />
    <span class="text-body-2">Go to {{ props.page }}</span>
    <div class="d-flex flex-column">
      <v-text-field
        placeholder=""
        variant="outlined"
        hide-details
        density="compact"
        type="number"
        :model-value="props.page"
        @update:model-value="changePageNumber"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { watch } from 'vue';

const props = defineProps<{
  itemsLength: number,
  itemsPerPage: number,
  page: number,
  pageCount: number
}>()

const emit = defineEmits([
  'update:itemsPerPage',
  'update:page'
])

const pageOptions = [
  { title: '5 / page', value: 5 },
  { title: '10 / page', value: 10 },
  { title: '20 / page', value: 20 },
  { title: '50 / page', value: 50 },
  { title: '100 / page', value: 100 },
]

const changePageNumber = (e: number | string) => {
  console.log('changePageNumber', e)
  if (isNaN(+e)) return;
  e = +e;
  if (e <= 0) return;
  if (e > props.pageCount) {
    e = props.pageCount;
  }
  emit('update:page', e)
}

watch(() => props.page, (e) => {
  emit('update:page', e)
})
</script>
