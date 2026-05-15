<script setup>
import { useDropdownLocation } from '@/composables/homePage/useDropdownLocation'
const { locationData } = useDropdownLocation()
</script>

<template>
  <div class="search-wrapper" id="search-dropdown-wrapper">
    <input class="nav-search" id="search-input" type="search" placeholder="🌏 Location">
    <span class="arrow-down"></span>
    <div class="search-menu" id="search-menu">
      <div class="menu-section locations">
        <h4>REGIONS</h4>
        <div class="menu-grid">
          <a v-for="item in locationData?.regions" :key="item.categoryItem" :href="item.url">
            {{ item.categoryItem }}
          </a>
        </div>
      </div>
      <hr class="menu-divider">
      <div class="menu-section locations">
        <h4>COUNTRIES</h4>
        <div class="menu-grid">
          <a v-for="item in locationData?.countries" :key="item.categoryItem" :href="item.url">
            {{ item.categoryItem }}
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>