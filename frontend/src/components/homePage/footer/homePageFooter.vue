<script setup>
import PropBtn from "@/components/globalBtns/propBtn.vue";
import { useFooter } from '@/composables/homePage/useFooter'
const { footer } = useFooter()
</script>

<template>
  <footer class="footer">
    <div class="container">
      <div class="footer-info">
        <p>{{ footer?.text }}</p>
        <div class="search-wrapper">
          <input class="nav-search" id="search-footer" :type="footer?.inputType" :placeholder="footer?.inputPlaceholder">
        </div>
        <PropBtn to="/post-job" text="Subscribe" className="btn btn-primary" id="post-button" />
      </div>
    </div>
    <div class="close-footer">{{ footer?.closeIcon }}</div>
    <div class="close-btn">
    </div>
  </footer>
</template>

<style scoped>

</style>