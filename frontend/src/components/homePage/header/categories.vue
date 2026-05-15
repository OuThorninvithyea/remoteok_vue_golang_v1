<script setup>
import { useCategories } from '@/composables/homePage/useCategories'

const { categories, loading, error } = useCategories()
</script>

<template>
  <div class="categories">
    <div class="category-wrapper" id="categories-container">
      <template v-if="loading">Loading...</template>
      <template v-else-if="error">{{ error }}</template>
      <template v-else>
        <template v-for="group in categories" :key="group.hoverText">
          <button v-for="filter in group.filters" :key="filter.label" class="category-button">
            <span class="hover-on-filter">{{ group.hoverText }}</span> {{ filter.emoji }} {{ filter.label }}
          </button>
        </template>
      </template>
    </div>
  </div>
</template>

<style scoped>

</style>