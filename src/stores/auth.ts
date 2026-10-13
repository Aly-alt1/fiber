// Utilities
import { UserAccount } from '../models/UserAccount'
import axios from 'axios'
import { defineStore } from 'pinia'
import { ref } from 'vue'

// axios.create({
//   withCredentials: true, // ¡importante para enviar cookies!
// })

export const useAuthStore = defineStore('auth', () => {
  const userAcc = ref<UserAccount>(new UserAccount())
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  const fetchProfile = async () => {
    isLoading.value = true
    try {
      userAcc.value = await axios.get('/api/current-user-acc')
    } catch (err) {
      userAcc.value = new UserAccount()
      error.value = 'No autorizado'
    } finally {
      isLoading.value = false
    }
  }

  return { userAcc, fetchProfile, isLoading, error }
})