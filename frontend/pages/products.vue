<template>
  <div>
    <h1>Products</h1>
    <div v-if="products.length">
      <ProductCard v-for="product in products" :key="product.id" :product="product" />
    </div>
    <p v-else>No products to show.</p>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import ProductCard from '~/components/ProductCard.vue'

const products = ref([])

onMounted(async () => {
  const { apiBase } = useRuntimeConfig().public
  try {
    const response = await fetch(`${apiBase}/products`)
    products.value = await response.json()
  } catch (error) {
    console.error('Error while fetching products:', error)
  }
})
</script>
