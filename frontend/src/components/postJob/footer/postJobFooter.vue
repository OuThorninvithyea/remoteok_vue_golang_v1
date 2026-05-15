<script setup>
import { usePostJobFooter } from '@/composables/postJob/footer/usePostJobFooter'
const { footerData } = usePostJobFooter()
</script>

<template>
    <footer class="post-job-footer">
        <div class="footer-form-zone">
            <div class="post-job-footer-preview">
                <img :src="footerData?.logo" alt="Logo" class="footer-company-logo">
                <div class="footer-job-info">
                    <p class="footer-company-name">{{ footerData?.companyPlaceholder }}</p>
                    <strong class="footer-position">{{ footerData?.positionPlaceholder }}</strong>
                    <span class="badge footer-badge">{{ footerData?.locationBadge }}</span>
                </div>
                <a href="#" class="btn preview-apply-btn footer-apply-btn">{{ footerData?.applyText }}</a>
            </div>
        </div>
        <div class="footer-sidebar-zone">
            <button class="btn btn-post-primary footer-post-btn">{{ footerData?.postButtonText }}</button>
        </div>
    </footer>
</template>

<style scoped>

</style>