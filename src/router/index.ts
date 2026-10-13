/**
 * router/index.ts
 *
 * Manual routes for ./src/pages/*.vue
 */

// Composables
import { createRouter, createWebHistory } from 'vue-router'
import Index from '@/pages/index.vue'
import Analytics from '@/pages/Analytics.vue'
import Settings from '@/pages/Settings.vue'
import Login from '@/pages/Login.vue'
import { useAuthStore } from '@/stores/auth'
import CardVuetify4 from '@/pages/CardVuetify4.vue'
import CardCafe from '@/pages/CardCafe.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: Index,
      meta: { requiresAuth: true },
    },
    {
      path: '/analytics',
      component: Analytics,
      meta: { requiresAuth: true },
    },
    {
      path: '/settings',
      component: Settings,
      meta: { requiresAuth: true },
    },
    {
      path: '/login',
      component: Login,
      meta: { requiresAuth: false },
    },
    {
      path: '/CardVuetify4',
      component: CardVuetify4,
      meta: { requiresAuth: false },
    },
    {
      path: '/cafe',
      component: CardCafe,
      meta: { requiresAuth: false },
    },
  ],
})

router.beforeEach(async (to, from, next) => {
  const authStore = useAuthStore()

  // Si no tenemos usuario, intentamos recuperarlo del backend (opcional)
  if (authStore.userAcc.id == '' && to.meta.requiresAuth) {
    console.log("Feching Profile ...")
    await authStore.fetchProfile() // Si la sesión es válida, retorna un usurio
  }

  if (to.meta.requiresAuth && authStore.userAcc.id == '') {
    console.log('redirecciona al /login')
    next('/login')
  } else if (to.meta.requiresGuest && authStore.userAcc.id != '') {
    console.log('redirecciona a raiz /')
    next('/')
  } else {
    console.log('redirecciona a pagina solicitada')
    next()
  }
})

export default router