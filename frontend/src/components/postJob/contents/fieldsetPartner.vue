<script setup>
import { usePostJobPartner } from '@/composables/postJob/form/usePostJobPartner'
const { partnerData } = usePostJobPartner()
</script>

<template>
    <legend>{{ partnerData?.legend }}</legend>
    <div class="health-insurance">
        <div class="logo-insurance">
            <img :src="partnerData?.logo" :alt="partnerData?.logoAlt">
        </div>
        <div class="job-details logo-title-insurance">
            <h2>{{ partnerData?.title }}</h2>
            <p>{{ partnerData?.description }}</p>
        </div>
        <a :href="partnerData?.buttonLink" class="btn sign-up health">{{ partnerData?.buttonText }}</a>
    </div>
</template>

<style scoped>

</style>