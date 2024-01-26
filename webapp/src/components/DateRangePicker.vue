<template>
  <span class="text-subtitle-1 font-weight-bold">{{ props.label ?? 'Date Range Picker' }}</span>
  <v-card
    class="d-flex flex-row align-center px-4"
    flat
    variant="outlined"
    :border="2"
    height="48px"
  >
    <v-menu :close-on-content-click="false">
      <template v-slot:activator="{ props }">
        <v-text-field
          placeholder="Start Date"
          variant="plain"
          hide-details
          prepend-inner-icon="mdi-calendar"
          clearable
          persistent-clear
          :model-value="startDateDisplay"
          density="comfortable"
          @click:clear="startDate = undefined"
          v-bind="props"
        ></v-text-field>
      </template>
      <template v-slot:default="{ isActive }">
        <v-date-picker
          color="primary"
          v-model="startDate"
          :max="endDate"
          hide-header
          @update:model-value="isActive.value = false"
        />
      </template>
    </v-menu>
    <span class="text-h5 mx-5">-</span>
    <v-menu :close-on-content-click="false">
      <template v-slot:activator="{ props }">
        <v-text-field
          placeholder="End Date"
          variant="plain"
          hide-details
          clearable
          persistent-clear
          :model-value="endDateDisplay"
          density="comfortable"
          @click:clear="endDate = undefined"
          v-bind="props"
        ></v-text-field>
      </template>
      <template v-slot:default="{ isActive }">
        <v-date-picker
          color="primary"
          v-model="endDate"
          :min="startDate"
          hide-header
          @update:model-value="isActive.value = false"
        />
      </template>
    </v-menu>
  </v-card>
</template>

<script setup lang="ts">
import dayjs from 'dayjs'
import { computed } from 'vue'

const props = defineProps<{
  label?: string
}>()
const startDate = defineModel<Date>('startDate')
const endDate = defineModel<Date>('endDate')

const startDateDisplay = computed(() => {
  return startDate.value ? dayjs(startDate.value).format('DD MMM YYYY') : ''
})

const endDateDisplay = computed(() => {
  return endDate.value ? dayjs(endDate.value).format('DD MMM YYYY') : ''
})
</script>
