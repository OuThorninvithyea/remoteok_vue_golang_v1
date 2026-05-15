<script setup>
import { useDropdownSearch } from '@/composables/homePage/useDropdownSearch'
const { searchData } = useDropdownSearch()
</script>

<template>
  <div class="search-wrapper" id="search-dropdown-wrapper">
    <input class="nav-search" id="search-input" type="search" placeholder="🔍 Search ">
    <span class="arrow-down"></span>
    <div class="search-menu" id="search-menu">
      <div class="menu-section">
        <a v-for="item in searchData" :key="item.CategoryItem" :href="item.URL">
          {{ item.categoryItem }}
        </a>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>