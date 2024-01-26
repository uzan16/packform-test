<template>
  <div data-testid="general-search" class="d-flex flex-row align-center">
    <v-icon size="x-large">mdi-magnify</v-icon>
    <span class="text-h5 ml-3 mr-6">Search</span>
    <v-text-field
      placeholder="Type here..."
      variant="outlined"
      hide-details="auto"
      density="comfortable"
      clearable
      persistent-clear
      v-model="model"
    ></v-text-field>
  </div>
</template>

<script setup lang="ts">
const model = defineModel()
</script>
