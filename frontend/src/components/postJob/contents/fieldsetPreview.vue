<script setup>
import { usePostJobPreview } from '@/composables/postJob/form/usePostJobPreview'
const { previewData } = usePostJobPreview()
</script>

<template>
    <legend>{{ previewData?.legend }}</legend>
    <div class="preview-info">
        <h2>{{ previewData?.heading }}</h2>
        <p>{{ previewData?.subheading }}</p>
    </div>
    <div class="preview-card">
        <div class="preview-card-header">
            <div class="logo-preview">
                <img :src="previewData?.card?.logo" alt="Logo">
                <div class="logo-name">
                    <p>{{ previewData?.card?.companyPlaceholder }}</p>
                    <h2>{{ previewData?.card?.positionPlaceholder }}</h2>
                    <span class="badge">{{ previewData?.card?.locationBadge }}</span>
                </div>
            </div>
            <a href="#" class="btn preview-apply-btn">Apply</a>
        </div>
        <div class="preview-text">
            <h1>{{ previewData?.card?.hiringText }}</h1>
            <h2>{{ previewData?.card?.roleText }}</h2>
        </div>
        <p>{{ previewData?.card?.descriptionPlaceholder }}</p>
        <div class="preview-card box">
            <h1>{{ previewData?.card?.applyHeading }}</h1>
            <p>{{ previewData?.card?.applyPlaceholder }}</p>
            <a href="#" class="btn preview-apply-btn preview">{{ previewData?.card?.applyButtonText }}</a>
        </div>
    </div>
    <div class="water-mark-preview">{{ previewData?.watermark }}</div>
</template>

<style scoped>

</style>