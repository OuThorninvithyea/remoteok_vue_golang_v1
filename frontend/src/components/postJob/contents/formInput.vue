<script setup>
import fieldsetLetStart from "@/components/postJob/contents/filedsetLetStart.vue"
import fieldsetDesignPost from "@/components/postJob/contents/fieldsetDesignPost.vue"
import fieldsetJobDetails from "@/components/postJob/contents/fieldsetJobDetails.vue"
import fieldsetCompany from "@/components/postJob/contents/fieldsetCompany.vue"
import fieldsetFeedback from "@/components/postJob/contents/fieldsetFeedback.vue"
import fieldsetPreview from "@/components/postJob/contents/fieldsetPreview.vue"
import fieldsetPartner from "@/components/postJob/contents/fieldsetPartner.vue"
</script>

<template>
    <fieldset>
        <fieldsetLetStart/>
    </fieldset>
    <fieldset>
        <fieldsetDesignPost/>
    </fieldset>
    <fieldset class="fieldset-3">
        <fieldsetJobDetails/>
    </fieldset>
    <fieldset>
        <fieldsetCompany/>
    </fieldset>
    <fieldset>
        <fieldsetFeedback/>
    </fieldset>
    <fieldset class="preview-fieldset">
        <fieldsetPreview/>
    </fieldset>
    <fieldset>
        <fieldsetPartner/>
    </fieldset>
</template>

<style scoped>

</style>
