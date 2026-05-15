<script setup>
import remoteLogoPostJob from "@/components/postJob/postJobBtns/remoteLogo.vue"
import PropBtn from "@/components/globalBtns/propBtn.vue"
import promotionPostJob from "@/components/postJob/header/promotionPostJob.vue"
import secPromotionPostJob from "@/components/postJob/header/secPromotionPostJob.vue"
import { usePostJobHeader } from '@/composables/postJob/header/usePostJobHeader'
const { headerData } = usePostJobHeader()
</script>

<template>
    <div class="post-job-header">
      <remoteLogoPostJob />
      <div class="post-job-header-title">
        <h1>{{ headerData?.Title }}</h1>
      </div>
        <PropBtn to="/post-job" text="Buy a bundle →" className="btn btn-primary" id="post-button" />
    </div>
  <promotionPostJob/>
  <secPromotionPostJob/>
</template>

<style scoped>

</style>