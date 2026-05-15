<script setup>
defineProps({
  job: { type: Object, required: true }
})
</script>

<template>
  <div class="job-card ad">
    <div class="job-card-main">
      <div class="job-info clickable-area">
        <div class="job-info-label">
          <img v-if="job.logo" :src="job.logo" alt="Company Logo" class="company-logo">
          <div class="job-details">
            <h3 v-if="job.title">{{ job.title }}</h3>
            <p>{{ job.company }}</p>
          </div>
        </div>
      </div>
      <div class="card-details-label">
        <p v-if="job.time" class="ad-text">{{ job.time }}</p>
        <a v-if="job.applyLabel" :href="job.applyLink || '#'" class="btn btn-primary">{{ job.applyLabel }}</a>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>
