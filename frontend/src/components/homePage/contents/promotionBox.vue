<script setup>
defineProps({
  job: { type: Object, required: true }
})
</script>

<template>
  <div class="promotion-container">
    <div v-if="job.img" class="promotion-img">
      <img :src="job.img" alt="promotonImage">
    </div>
    <div class="promotion-text">
      <h2 v-if="job.title" class="promotion-title">{{ job.title }}</h2>
      <div v-if="job.paragraph" class="promotion-paragraph">
        <p>{{ job.paragraph }}</p>
      </div>
      <div v-if="job.rating" class="promotion-rating">{{ job.rating }}</div>
      <div v-if="job.button" class="get-acces-btn btn btn-primary">{{ job.button }}</div>
    </div>
  </div>
</template>

<style scoped>

</style>
