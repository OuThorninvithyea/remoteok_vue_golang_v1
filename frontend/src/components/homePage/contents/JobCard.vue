<script setup>
import AdCard from "@/components/homePage/contents/adCard.vue"
import VerifiedCard from "@/components/homePage/contents/verifiedCard.vue"
import HighlightCard from "@/components/homePage/contents/highlightCard.vue"
import NormalCard from "@/components/homePage/contents/normalCard.vue"
import PromotionBox from "@/components/homePage/contents/promotionBox.vue"

defineProps({
  job: { type: Object, required: true }
})
</script>

<template>
  <AdCard        v-if="job.type === 'ad'"   :job="job" />
  <VerifiedCard  v-else-if="job.type === 'verified'"  :job="job" />
  <HighlightCard v-else-if="job.type === 'highlight'" :job="job" />
  <PromotionBox  v-else-if="job.type === 'promotion'" :job="job" />
  <NormalCard    v-else                               :job="job" />
</template>
