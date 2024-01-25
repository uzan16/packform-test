<script setup lang="ts">
import { RouterView } from 'vue-router'
</script>

<template>
  <main>
    <v-sheet
      class="mx-auto py-8 px-10 my-4"
      :rounded="true"
      :border="true"
      width="95vw"
    >
      <RouterView />
    </v-sheet>
  </main>
</template>
