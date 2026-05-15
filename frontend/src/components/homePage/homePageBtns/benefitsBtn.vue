
<script setup>
  import { useDropdownBenefits } from '@/composables/homePage/useDropdownBenefits'
  const { benefitsData } = useDropdownBenefits()
</script>

<template>
  <div class="search-wrapper" id="benefits-dropdown-wrapper">
    <input class="nav-search" id="benefits-input" type="search" placeholder="🎪 Benefits">
    <span class="arrow-down"></span>
    <div class="search-menu" id="benefits-menu">
      <div class="menu-section">
        <a v-for="item in benefitsData" :key="item.categoryItem" :href="item.url">
          {{ item.categoryItem }}
        </a>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>