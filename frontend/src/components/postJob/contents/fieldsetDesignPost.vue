<script setup>
import { usePostJobDesignYourJob } from '@/composables/postJob/form/usePostJobDesignYourJob'
const { designYourJobData } = usePostJobDesignYourJob()
</script>

<template>
    <legend>{{ designYourJobData?.legend }}</legend>
    <div class="check-box-post-job">
        <label v-for="checkbox in designYourJobData?.checkboxes" :key="checkbox.name" class="custom-checkbox-container">
            <input type="checkbox" :name="checkbox.name" class="hidden-checkbox" :checked="checkbox.checked">
            <span class="custom-box"></span>
            <span class="checkbox-text">
                {{ checkbox.text }}
                <a v-if="checkbox.linkText" :href="checkbox.linkUrl">{{ checkbox.linkText }}</a>
                <span v-if="checkbox.afterLinkText">{{ checkbox.afterLinkText }}</span>
            </span>
            <span v-if="checkbox.badge" class="btn btn-secondary check-box">{{ checkbox.badge }}</span>
        </label>
        <div class="view-counter">
            <div class="expected-view-counter">
                Expected views <span><label>
                    <input class="position-input" type="text" :placeholder="designYourJobData?.viewCounterPlaceholder">
                </label></span>
            </div>
            <div class="expected-clicks-counter">
                Expected clicks <span><label>
                    <input class="position-input" type="text" :placeholder="designYourJobData?.clicksCounterPlaceholder">
                </label></span>
            </div>
        </div>
    </div>
    <p>{{ designYourJobData?.disclaimer }}</p>
</template>

<style scoped>

</style>