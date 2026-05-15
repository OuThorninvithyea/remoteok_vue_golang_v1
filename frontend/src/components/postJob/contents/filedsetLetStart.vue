<script setup>
import { usePostJobLetStart } from '@/composables/postJob/form/usePostJobLetStart'
import { usePostJobToolbar } from '@/composables/postJob/form/usePostJobToolbar'
const { letStartData, loading, error } = usePostJobLetStart()
const { toolbarData } = usePostJobToolbar()
</script>

<template>
    <legend>{{ letStartData?.legend }}</legend>
    <div class="position-details">
      <label>{{ letStartData?.position?.label }}</label>
      <label>
        <input type="text" :placeholder="letStartData?.position?.placeholder" class="position-input" :required="letStartData?.position?.required">
      </label>
      <p>{{ letStartData?.position?.helpText }}</p>
    </div>
    <div class="company-name-details">
      <label>{{ letStartData?.companyName?.label }}</label>
      <label>
        <input type="text" :placeholder="letStartData?.companyName?.placeholder" class="position-input" :required="letStartData?.companyName?.required">
      </label>
      <p>{{ letStartData?.companyName?.helpText }}</p>
    </div>
    <div class="job-description-input">
      <label>{{ letStartData?.jobDescription?.label }} <span class="write-with-ai btn btn-post-primary">{{ letStartData?.jobDescription?.aiButtonText }}</span>
      </label>
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
    <div class="employee-type-details">
      <label for="employment-type">EMPLOYMENT TYPE*</label>
      <select id="employment-type" class="position-input interactive">
        <option v-for="type in letStartData?.employmentTypes" :key="type.value" :value="type.value">{{ type.label }}</option>
      </select>
    </div>
    <div class="primary-tag-details">
      <label for="primary-tag">PRIMARY TAG</label>
      <select id="primary-tag" class="position-input interactive">
        <option v-for="tag in letStartData?.primaryTags" :key="tag.value" :value="tag.value" :disabled="tag.disabled" :selected="tag.disabled">{{ tag.label }}</option>
      </select>
      <p>{{ letStartData?.primaryTagHelpText }}</p>
    </div>
    <div class="keyword-tags-details">
      <label for="tag-input">{{ letStartData?.tags?.label }}</label>
      <div class="tag-input-wrapper">
        <div class="tag-selected-list" id="tag-selected-list"></div>
        <input id="tag-input" type="text" :placeholder="letStartData?.tags?.placeholder" class="position-input tag-search-input" autocomplete="off">
        <ul class="tag-dropdown" id="tag-dropdown"></ul>
      </div>
      <p>{{ letStartData?.tags?.helpText }}</p>
    </div>
    <div class="job-location-details">
      <label class="tags-job">{{ letStartData?.location?.label }} <span v-if="letStartData?.location?.isNew" class="tags btn-primary">new</span></label>
      <label>
        <input type="text" :placeholder="letStartData?.location?.placeholder" class="position-input" :required="letStartData?.location?.required">
      </label>
      <p>{{ letStartData?.location?.helpText }} <span class="job-locations-p highlight-text">{{ letStartData?.location?.highlightText }}</span>  <span class="text-underline">{{ letStartData?.location?.underlineText }}</span></p>
    </div>
</template>

<style scoped>

</style>
