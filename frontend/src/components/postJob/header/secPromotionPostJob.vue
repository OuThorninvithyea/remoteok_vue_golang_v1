<script setup>
import { usePostJobPromotion } from '@/composables/postJob/header/usePostJobPromotion'
const { promotionData } = usePostJobPromotion()
</script>

<template>
  <div class="sec-promotion">
    {{ promotionData?.secPromotion?.icon }} {{ promotionData?.secPromotion?.text }}
  </div>
</template>

<style scoped>

</style>
