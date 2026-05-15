<script setup>
import { useDropdownSalary } from '@/composables/homePage/useDropdownSalary'
const { salaryData } = useDropdownSalary()
</script>

<template>
  <div class="search-wrapper" id="salary-dropdown-wrapper">
    <input class="nav-search" id="salary-input" type="search" placeholder="💵 Salary">
    <span class="arrow-down"></span>
    <div class="search-menu salary-menu" id="salary-menu">
      <div class="salary-header">
        <span class="salary-title">Minimum</span>
        <span class="salary-value" id="salary-value-display">{{ salaryData?.currencyPrefix }}{{ salaryData?.value }}{{ salaryData?.currencySuffix }}</span>
      </div>
      <div class="salary-slider-container">
        <input type="range" id="salary-slider" :min="salaryData?.min" :max="salaryData?.max" :step="salaryData?.step" :value="salaryData?.value" class="salary-slider">
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>