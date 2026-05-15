<script setup>
import { usePostJobFeedback } from '@/composables/postJob/form/usePostJobFeedback'
const { feedbackData } = usePostJobFeedback()
</script>

<template>
    <legend>{{ feedbackData?.legend }}</legend>
    <div class="position-details feed-back">
        <label>{{ feedbackData?.label }}</label>
        <label>
            <textarea :placeholder="feedbackData?.placeholder" class="position-input textarea-large"></textarea>
        </label>
        <p>{{ feedbackData?.helpText }}</p>
    </div>
</template>

<style scoped>

</style>