<template>
  <div>
    <h1>Here is my site!</h1>
    <p>Status API: {{ apiStatus }}</p>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const apiStatus = ref('')

onMounted(async () => {
  const { apiBase } = useRuntimeConfig().public
  try {
    const response = await fetch(`${apiBase}/health`)
    const data = await response.json()
    apiStatus.value = data.status
  } catch (error) {
    apiStatus.value = 'Error connection with API'
  }
})
</script>
