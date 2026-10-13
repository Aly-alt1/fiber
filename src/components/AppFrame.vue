<template>
  <v-app>
    <v-navigation-drawer v-model="drawer">
      <template v-slot:prepend>
          <v-list-item
            lines="two"
            prepend-avatar="https://cdn.vuetifyjs.com/docs/images/team/john.png"
            subtitle="Logged in"
            title="John Leider"
          ></v-list-item>
        </template>

        <v-divider></v-divider>
        <v-list>
          <v-list-item v-for="(mnu, idx) in menuItems" :key="idx" :prepend-icon="mnu.icon" :title="mnu.text" :to="mnu.to"></v-list-item>
        </v-list>
    </v-navigation-drawer>

    <v-app-bar color="primary">
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>

      <v-app-bar-title>Cummins nuevos servicios</v-app-bar-title>

      <template v-slot:append>
        <v-btn icon="mdi-theme-light-dark"></v-btn>
        <v-btn icon="mdi-translate"></v-btn>
        <v-btn icon="mdi-exit-to-app" @click="logout"></v-btn>
      </template>
    </v-app-bar>

    <v-main>
      <v-slide-x-reverse-transition mode="out-in" group>
        <v-container fluid>
          <slot></slot>
        </v-container>
      </v-slide-x-reverse-transition>
    </v-main>
  </v-app>
</template>

<script lang="ts" setup>
import axios from 'axios';
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router';
import { useTheme } from 'vuetify'

const router = useRouter()
const drawer = ref(true)
const theme = useTheme()

const menuItems = [
  { icon: 'mdi-view-dashboard-outline', text:'Dashboard', to: '/'},
  { icon: 'mdi-chart-line', text: 'Analytics', to: '/analytics' },
  { icon: 'mdi-cog-outline', text: 'Settings', to: '/settings'},
  //{ icon: 'mdi-login', text: 'Login', to: '/login'},
]

// onMounted(() => {
//   axios.get('/api/profiles/x')
//   .then(response => response.data)
//   .catch(error => router.push('/login'))
// })

const logout = () => {
  axios.post('/api/logout')
  .then(response => {
    if (response.data == 'success') {
      router.push('/login')
    }
  })
}
</script>