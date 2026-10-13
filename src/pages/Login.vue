<template>
  <v-container class="h-100 pa-4" fluid>
    <v-row class="h-100 align-center justify-center">
      <v-sheet class="flex-1-1 px-4" color="transparent" max-width="420">
        <v-img class="mx-auto mb-4" height="60" max-width="60" src="https://cdn.vuetifyjs.com/docs/images/logos/v.svg"/>

        <div class="text-headline-small text-center mb-8 font-weight-medium">Registrarse con su cuenta</div>

        <v-sheet border="opacity-25 thin" class="overflow-hidden" rounded="lg">
          <v-text-field v-model="credentials.username" placeholder="Nombre de usuario" class="ma-3" rounded="lg" variant="solo" hide-details/>
          <v-divider></v-divider>
          <v-text-field v-model="credentials.password" placeholder="Contraseña" type="password" class="ma-3" variant="solo" hide-details/>
          <v-divider></v-divider>
          <v-text-field v-model="credentials.role"  placeholder="Role" rounded="lg" class="ma-3" variant="solo" hide-details/>
        </v-sheet>

        <div class="py-4">
          <div class="d-flex justify-space-between align-center">
            <v-checkbox-btn class="ms-n3" color="primary" label="Remember me">
              <template #label>
                <span class="text-body-medium">Remember me</span>
              </template>
            </v-checkbox-btn>

            <a class="text-decoration-none text-primary text-body-medium font-weight-medium cursor-pointer" @click="login">
              Forgot password?
            </a>
          </div>
        </div>

        <v-btn class="text-none mb-8" color="primary" rounded="lg" text="Log In" block flat @click="login"/>

        <div class="text-center text-body-medium">
          Don't have an account?
          <a class="text-decoration-none text-primary font-weight-medium cursor-pointer" @click="login">
            Sign up
          </a>
        </div>
      </v-sheet>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import axios from 'axios'
import { defineComponent, reactive } from 'vue'
import { useRouter } from 'vue-router'

export default defineComponent({
  name: 'Login',
  components: {},
   
  setup() {
    const router = useRouter()
    const credentials = reactive({ username: '', password: '', role: ''})

    const login = () => {
      axios.post("/api/login", credentials)
      .then(response => {
        if (response.data == 'success') {
          router.push('/')
        }
      })
    }

    return { credentials, login }
  }
})
</script>