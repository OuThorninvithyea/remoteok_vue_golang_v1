<script setup>
import { usePostJobJobDetails } from '@/composables/postJob/form/usePostJobJobDetails'
import { usePostJobToolbar } from '@/composables/postJob/form/usePostJobToolbar'
const { jobDetailsData } = usePostJobJobDetails()
const { toolbarData } = usePostJobToolbar()
</script>

<template>
    <legend>{{ jobDetailsData?.legend }}</legend>

    <div class="position-details img">
        <label>{{ jobDetailsData?.logoUpload?.label }}</label>
        <label for="company-logo" class="upload-btn"><span>{{ jobDetailsData?.logoUpload?.buttonText }}</span></label>
        <input type="file" id="company-logo" class="upload-image">
    </div>

    <div class="position-details">
        <label class="custom-checkbox-container">
            <input type="checkbox" :name="jobDetailsData?.brandColor?.name" class="hidden-checkbox" :checked="jobDetailsData?.brandColor?.checked">
            <span class="custom-box"></span>
            <span class="checkbox-text">{{ jobDetailsData?.brandColor?.text }} <span class="color-swatch"></span></span>
            <span class="btn btn-secondary show-rate">{{ jobDetailsData?.brandColor?.badge }}</span>
        </label>
    </div>

    <div class="position-details">
        <label>{{ jobDetailsData?.salary?.label }}</label>
        <div class="salary-range">
            <label>
                <input type="number" :placeholder="jobDetailsData?.salary?.minPlaceholder" :min="jobDetailsData?.salary?.min" :max="jobDetailsData?.salary?.max" :step="jobDetailsData?.salary?.step" class="position-input minimum-salary" required>
            </label>
            <span class="salary-dash">—</span>
            <label>
                <input type="number" :placeholder="jobDetailsData?.salary?.maxPlaceholder" :min="jobDetailsData?.salary?.min" :max="jobDetailsData?.salary?.max" :step="jobDetailsData?.salary?.step" class="position-input maximun-salary" required>
            </label>
        </div>
        <p>{{ jobDetailsData?.salary?.helpText }}</p>
    </div>

    <div class="position-details benefits">
        <label class="tags-job">BENEFITS<span class="tags btn-primary">new</span></label>
        <div class="benefits-container">
            <span v-for="benefit in jobDetailsData?.benefits" :key="benefit">{{ benefit }}</span>
        </div>
    </div>

    <div class="job-description-input how-to-apply">
        <label>HOW TO APPLY</label>
        <div class="form-description position-container">
            <div class="form-nav">
                <div v-for="toolbar in toolbarData" :key="toolbar.group" :class="toolbar.group">
                    <img v-for="icon in toolbar.icons" :key="icon" class="form-nav-img" :src="icon" :alt="toolbar.group">
                </div>
            </div>
            <label>
                <input class="input-button-form">
            </label>
        </div>
    </div>

    <div class="position-details apply-url">
        <label>{{ jobDetailsData?.applyUrl?.label }}</label>
        <label>
            <input type="text" :placeholder="jobDetailsData?.applyUrl?.placeholder" class="position-input" :required="jobDetailsData?.applyUrl?.required">
        </label>
        <p>{{ jobDetailsData?.applyUrl?.helpText }}</p>
    </div>

    <div class="position-details apply-email">
        <label>{{ jobDetailsData?.applyEmail?.label }}</label>
        <label>
            <input type="text" :placeholder="jobDetailsData?.applyEmail?.placeholder" class="position-input" :required="jobDetailsData?.applyEmail?.required">
        </label>
        <p>{{ jobDetailsData?.applyEmail?.helpText }}</p>
    </div>
</template>

<style scoped>

</style>