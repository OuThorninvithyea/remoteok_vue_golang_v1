<script setup>
import HeaderBtn from "@/components/homePage/header/header.vue"
import MainBar from "@/components/homePage/header/mainBar.vue"
import NavigationBar from "@/components/homePage/header/navigationBar.vue"
import JobCard from "@/components/homePage/contents/JobCard.vue"
import HomePageFooter from "@/components/homePage/footer/homePageFooter.vue"
import { useJobs } from '@/composables/homePage/useJobs'
const { jobs } = useJobs()
</script>

<template>
  <header>
    <div class="header">
      <nav class="nav-links">
        <HeaderBtn/>
      </nav>
    </div>
  </header>
  <main>
    <div class="home-page">
      <div class="main-bar">
        <MainBar/>
      </div>
    </div>
    <NavigationBar/>
  </main>
  <section class="job-listing">
    <div class="container" id="job-listings-container">
      <JobCard v-for="(job, index) in jobs" :key="index" v-bind="{ job }"/>
    </div>
  </section>
  <footer class="footer">
  <HomePageFooter/>
  </footer>
</template>

<style scoped>
  .container {
    margin-top: 10px;
  }
</style>
