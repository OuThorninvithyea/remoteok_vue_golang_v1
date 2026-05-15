<script setup>
import { useTrustedCompanies } from '@/composables/postJob/sidebar/useTrustedCompanies'
const { trustedCompaniesData } = useTrustedCompanies()
</script>

<template>
    <div class="trusted-companies-header">
        <span class="text-underline highlight-text text-marked">{{ trustedCompaniesData?.highlightText }}</span>
        {{ trustedCompaniesData?.headerText }}
    </div>

    <div class="trusted-companies-logos">
        <img v-for="logo in trustedCompaniesData?.logos" :key="logo.alt" :src="logo.src" :alt="logo.alt">
        <p>{{ trustedCompaniesData?.showMoreText }}</p>
    </div>
</template>

<style scoped>

</style>