<script setup>
import PropBtn from "@/components/globalBtns/propBtn.vue";
import { useMainBar } from '@/composables/homePage/useMainBar'
const { mainBar } = useMainBar()
</script>

<template>
  <div class="main-bar">
    <div class="bar-container">
      <p>{{ mainBar?.emoji }} {{ mainBar?.text }}
        <a :href="mainBar?.highlightLink" class="highlight-text text-underline bar">{{ mainBar?.highlightText }}</a>
        {{ mainBar?.suffix }}
      </p>
      <div class="bar-container-button">
        <nav class="nav-links">
          <PropBtn to="/post-job" text="Post a remote job →" className="btn btn-primary" id="post-button" />
          <PropBtn to="/" text="Hide this" className="btn btn-secondary" id="post-button" />
        </nav>
      </div>
    </div>
  </div>
</template>

<style scoped>


</style>

