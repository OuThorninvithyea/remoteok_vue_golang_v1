import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('@/views/LoginView.vue'),
  },
  {
    path: '/post-job',
    name: 'postJob',
    component: () => import('@/views/PostJobView.vue'),
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router