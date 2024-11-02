<template>
  <div class="product-card">
    <h2>{{ product.name }}</h2>
    <p>{{ product.description }}</p>
    <p>Cena: {{ product.price }} PLN</p>
  </div>
</template>

<script setup>
defineProps({
  product: {
    type: Object,
    required: true
  }
})
</script>

<style scoped>
.product-card {
  border: 1px solid #ddd;
  padding: 1rem;
  margin: 0.5rem 0;
}
</style>
