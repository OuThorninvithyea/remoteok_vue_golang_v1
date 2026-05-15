<script setup>
import { useLogo } from '@/composables/homePage/useLogo'
const { logo } = useLogo();
</script>

<template>
  <div class="logo-dropdown" id="logo-click-dropdown">
    <button class="logo-btn" id="logo-toggle-btn">
      <img src="/images/image_003_2966.png" alt="Remote OK" class="logo-img">
      <span class="arrow-down"></span>
    </button>
    <div class="logo-menu" id="logo-dropdown-menu">
      <div class="menu-section header-logo-clicked-info">
        <a href="#">{{ logo?.header?.join }}</a>
        <a href="#">{{ logo?.header?.login }}</a>
      </div>
      <template v-for="(section, i) in logo?.sections" :key="i">
        <hr class="menu-divider">
        <div class="menu-section">
          <h4>{{ section.title }}</h4>
          <div class="menu-grid">
            <a v-for="item in section.items" :key="item.categoryItem" :href="item.url">
              {{ item.categoryItem }}
            </a>
          </div>
        </div>
      </template>
    </div>
  </div>
</template>

<style scoped>

</style>