<script setup>
  defineProps({
    to: {type: String, default: "/"},
    text: {type: String, required: true},
    className: {type: String, default: ''},
    id: {type: String, default: null},
  })
</script>

<template>
  <router-link :to="to" :class="className" v-bind:id="id">{{text}}</router-link>
</template>

<style scoped>

</style>