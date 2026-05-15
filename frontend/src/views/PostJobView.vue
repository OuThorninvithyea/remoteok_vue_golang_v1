<script setup>
import PostJobHeader from '@/components/postJob/header/postJobHeader.vue'
import formInput from '@/components/postJob/contents/formInput.vue'
import sideBar from '@/components/postJob/contents/sideBar.vue'
import postJobFooter from '@/components/postJob/footer/postJobFooter.vue'
import '@/styles/postJob.css'
</script>

<template>
  <header>
    <PostJobHeader />
  </header>
  <main>
    <div class="main">
      <div class="form-input">
        <form class="form-container">
          <formInput/>
        </form>
      </div>
      <sideBar/>
    </div>
  </main>
  <postJobFooter/>
</template>

<style scoped>
</style>