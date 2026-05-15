<script setup>
import logoDropDownBtn from "@/components/homePage/homePageBtns/logoBtn.vue";
import seracBtn from "@/components/homePage/homePageBtns/searchBtn.vue";
import locationBtn from "@/components/homePage/homePageBtns/locationBtn.vue";
import salaryBtn from "@/components/homePage/homePageBtns/salaryBtn.vue";
import benefitsBtn from "@/components/homePage/homePageBtns/benefitsBtn.vue";
import sortbyBtn from "@/components/homePage/homePageBtns/sortbyBtn.vue";
import categories from "@/components/homePage/header/categories.vue"
</script>

<template>
  <section class="navigation-bar">
    <div class="sec-header">
      <div class="navs">
        <div class="search-group">
          <logoDropDownBtn/>
          <seracBtn/>
          <locationBtn/>
          <salaryBtn/>
          <benefitsBtn/>
        </div>
          <sortbyBtn/>
      </div>
      <categories/>
    </div>
  </section>
</template>

<style scoped>

</style>
