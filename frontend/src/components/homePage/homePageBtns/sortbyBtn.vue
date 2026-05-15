<script setup>
import { useDropdownSort } from '@/composables/homePage/useDropdownSort'
const { sortData } = useDropdownSort()
</script>

<template>
  <div class="search-wrapper" id="sort-filter">
    <input class="nav-search" id="sort-input" type="search" placeholder="🦴 Sort by">
    <span class="arrow-down"></span>
    <div class="search-menu sort-menu" id="sort-menu">
      <a v-for="item in sortData" :key="item.categoryItem" :href="item.url">{{ item.categoryItem }}</a>
    </div>
  </div>
</template>

<style scoped>

</style>