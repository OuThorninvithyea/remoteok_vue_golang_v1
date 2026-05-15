<script setup>
import { usePostJobPromotion } from '@/composables/postJob/header/usePostJobPromotion'
const { promotionData } = usePostJobPromotion()
</script>

<template>
  <div class="promotion">
    {{ promotionData?.promotion?.icon }} {{ promotionData?.promotion?.text }}
  </div>
</template>

<style scoped>

</style>
