<script setup>
import PropBtn from "@/components/globalBtns/propBtn.vue";
</script>

<template>
  <nav class="nav-links">
    <a href="#" class="health" id="health-button">Health insurance</a>
    <PropBtn to="/post-job" text="Post a job →" className="btn btn-primary" id="post-button" />
    <PropBtn to="/login" text="Log in" className="btn btn-secondary" id="login-button" />
  </nav>
</template>

<style scoped>

</style>