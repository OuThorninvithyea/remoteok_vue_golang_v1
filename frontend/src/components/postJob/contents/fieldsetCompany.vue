<script setup>
import { usePostJobCompany } from '@/composables/postJob/form/usePostJobCompany'
const { companyData } = usePostJobCompany()
</script>

<template>
    <legend>{{ companyData?.legend }}</legend>

    <div v-for="field in companyData?.fields" :key="field.label" class="position-details">
        <label>{{ field.label }}</label>
        <label>
            <textarea v-if="field.type === 'textarea'" :placeholder="field.placeholder" :class="['position-input', field.textareaClass]"></textarea>
            <input v-else :type="field.type" :placeholder="field.placeholder" class="position-input" :required="field.required">
        </label>
        <p>{{ field.helpText }}</p>
    </div>

    <label>{{ companyData?.payLaterLabel }}</label>

    <label class="custom-checkbox-container">
        <input type="checkbox" :name="companyData?.showLogoCheckbox?.name" class="hidden-checkbox" :checked="companyData?.showLogoCheckbox?.checked">
        <span class="custom-box"></span>
        <span class="checkbox-text">{{ companyData?.showLogoCheckbox?.text }}</span>
    </label>
</template>

<style scoped>

</style>