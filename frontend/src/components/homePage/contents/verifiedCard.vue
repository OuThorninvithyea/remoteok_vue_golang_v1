<script setup>
import { ref } from "vue"
import ExpendCard from "@/components/globalFuncExpendCard/expendCard.vue"

defineProps({
  job: { type: Object, required: true }
})

const isExpanded = ref(false)
</script>

<template>
  <div :class="['job-card', 'verified', { 'expanded': isExpanded }, job.cardClass || '']">
    <div class="job-card-main" @click="isExpanded = !isExpanded" style="cursor: pointer;">
      <div class="job-info clickable-area">
        <div class="job-info-label">
          <img v-if="job.logo" :src="job.logo" alt="Company Logo" class="company-logo">
          <div class="job-details">
            <div class="title-verified">
              <h3 v-if="job.title">{{ job.title }}</h3>
              <span v-if="job.badgeClass" :class="[job.badgeClass, 'verified-label']">VERIFIED</span>
            </div>
            <p>
              {{ job.company }}
              <span v-if="job.isNew" class="new-gif-stlye">
                <img src="/images/new2x.gif" alt="New">
              </span>
            </p>
            <div class="job-badges">
              <span v-if="job.locationBadge" class="badge">{{ job.locationBadge }}</span>
              <span v-if="job.salaryBadge" class="badge">{{ job.salaryBadge }}</span>
              <span v-if="job.typeBadge" class="badge">{{ job.typeBadge }}</span>
            </div>
          </div>
        </div>
        <div class="job-badges">
          <span v-for="badge in job.middleBadges" :key="badge" class="badge middle-badge black-border">
            <span class="hover-on-filter">Add to filters</span>
            {{ badge }}
          </span>
          <span v-for="badge in job.extraBadges" :key="badge" class="badge extra-badge">
            {{ badge }}
          </span>
        </div>
      </div>
      <div class="card-details-label">
        <p v-if="job.time" class="ad-text">{{ job.time }}</p>
        <a v-if="job.applyLabel" :href="job.applyLink || '#'" class="btn btn-primary">{{ job.applyLabel }}</a>
      </div>
    </div>

    <ExpendCard :job="job" />
  </div>
</template>

<style scoped>

</style>
