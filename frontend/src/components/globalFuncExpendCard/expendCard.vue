<script setup>
defineProps({
  job: { type: Object, required: true }
})
</script>

<template>
  <div v-if="job.detail" class="job-details-expanded">
    <div class="expanded-job-details">
      <div class="description-column">
        <p v-if="job.detail.hiringLabel" class="hiring-label">{{ job.detail.hiringLabel }}</p>
        <h2 v-if="job.detail.roleTitle" class="expanded-title">{{ job.detail.roleTitle }}</h2>

        <div class="expanded-paragraphs">
          <p v-for="paragraph in job.detail.paragraphs" :key="paragraph">{{ paragraph }}</p>
        </div>

        <div class="expanded-list-section">
          <h3>Your Responsibilities</h3>
          <ul>
            <li v-for="item in job.detail.responsibilities" :key="item">{{ item }}</li>
          </ul>
        </div>
      </div>

      <div class="sidebar-container">
        <div class="sidebar-card">
          <img v-if="job.logo" :src="job.logo" :alt="job.detail.companyName || ''" class="sidebar-logo">
          <div class="sidebar-info">
            <h3 v-if="job.detail.text" class="sidebar-text">{{ job.detail.text }}</h3>
            <p v-if="job.detail.applyUrl">{{ job.detail.applyUrl }}</p>
            <a href="#" class="btn btn-secondary btn-expanded-apply reve">Apply now</a>
            <p class="share-this-job">Share this job:</p>
            <div class="qr-code-section">
              <img v-if="job.detail.qrCode" :src="job.detail.qrCode" alt="QR Code" class="qr-img">
              <a href="#" class="btn-secondary btn rok">rok.co/ri</a>
              <p class="qr-text">Also hiring:</p>
              <p class="hire">Entry level Crypto Market...</p>
            </div>
          </div>
        </div>
      </div>

      <div class="expanded-list-section offer">
        <h3>What we offer</h3>
        <ul>
          <li v-for="item in job.detail.offers" :key="item">{{ item }}</li>
        </ul>
      </div>

      <div class="expanded-list-section interview-process">
        <h3>Interview Process</h3>
        <ul>
          <li v-for="step in job.detail.interviewProcess" :key="step">{{ step }}</li>
        </ul>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>
