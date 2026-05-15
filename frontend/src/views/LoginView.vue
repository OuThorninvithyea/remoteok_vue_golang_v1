<script setup>
import LogoBtn from '@/components/homePage/homePageBtns/logoBtn.vue'
import HeaderBtn from '@/components/homePage/header/header.vue'
import LoginForm from '@/components/login/loginForm.vue'
import { useLogin } from '@/composables/login/useLogin'
const { loginData } = useLogin()
</script>

<template>
  <div class="background-img-login">
    <img :src="loginData?.backgroundImg" alt="Promotion">
  </div>
  <header>
    <div class="header">
      <LogoBtn />
      <HeaderBtn />
    </div>
  </header>
  <main>
    <LoginForm />
  </main>

  <footer>
  </footer>

</template>

<style>
@import '@/styles/login.css';
</style>