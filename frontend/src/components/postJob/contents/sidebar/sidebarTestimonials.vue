<script setup>
import { useSidebarTestimonials } from '@/composables/postJob/sidebar/useSidebarTestimonials'
const { testimonialsData, founderData } = useSidebarTestimonials()
</script>

<template>
    <template v-for="testimonial in testimonialsData" :key="testimonial.name">
        <div class="border-bottom"></div>
        <div class="sidebar-stat profile">
            <img :src="testimonial.image" :alt="testimonial.alt">
            <h3>" <span v-if="testimonial.beforeHighlight" class="text-muted">{{ testimonial.beforeHighlight }} </span><span class="text-bg-yellow">{{ testimonial.highlight }}</span><span v-if="testimonial.muted" class="text-muted"> {{ testimonial.muted }}</span></h3>
            <h3>— <span class="text-underline">{{ testimonial.name }}</span>, {{ testimonial.role }} at <span class="text-underline">{{ testimonial.company }}</span></h3>
        </div>
    </template>

    <div class="border-bottom"></div>
    <div class="sidebar-stat profile peter">
        <img :src="founderData?.image" :alt="founderData?.alt">
        <p v-for="paragraph in founderData?.bio" :key="paragraph" class="peter-bio" v-html="paragraph"></p>
        <h3>— <span class="text-underline">{{ founderData?.name }}</span>, {{ founderData?.role }} of <span class="text-underline">{{ founderData?.company }}</span></h3>
    </div>
    <div class="border-bottom"></div>
</template>

<style scoped>

</style>