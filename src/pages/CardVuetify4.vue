<template>
  <v-container>
    <v-card class="mx-auto pa-4" max-width="500" elevation="2" rounded="lg">
      <v-card-item class="pa-0 mb-4">
        <template v-slot:append>
          <v-btn icon="mdi-dots-horizontal" variant="text" size="small"></v-btn>
        </template>
        <v-card-title class="text-subtitle-1 font-weight-bold">
          Active projects
        </v-card-title>
      </v-card-item>

      <div class="d-flex mb-6" style="gap: 4px;">
        <div
          v-for="(segment, index) in segments" :key="index"
          :style="{
            backgroundColor: segment.color,
            height: '24px',
            flex: '1',
            borderRadius: '2px'
          }">
        </div>
      </div>

      <v-row dense>
        <v-col v-for="item in statusData" :key="item.label" cols="6">
          <v-sheet
            color="grey-lighten-4"
            class="pa-3 d-flex align-center justify-space-between rounded-lg"
          >
            <div class="d-flex align-center">
              <v-badge :color="item.color" dot inline class="mr-2"></v-badge>
              <span class="text-body-2 font-weight-medium">{{ item.label }}</span>
            </div>
            <v-chip size="x-small" variant="flat" color="grey-lighten-2" class="text-caption">
              {{ item.value }}%
            </v-chip>
          </v-sheet>
        </v-col>
      </v-row>
    </v-card>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue'

export default defineComponent({
   name: 'CardVuetify4',
   components: {},

   setup() {
      const statusData = [
         { label: 'Backlog', value: 30, color: '#26C6DA' },
         { label: 'Design', value: 40, color: '#2196F3' },
         { label: 'New', value: 25, color: '#E57373' },
         { label: 'Development', value: 5, color: '#7CB342'},
      ]
      // Generamos 40 segmentos totales para que se vea igual a la imagen
      const segments = computed(() => {
         const totalBars = 40
         //const result = []
         
         statusData.forEach(status => {
            const count = Math.round((status.value / 100) * totalBars)
            for (let i = 0; i < count; i++) {
               //result.push({ color: status.color })
            }
         })

         return //result
      })

      return { statusData, segments }
   }
})
</script>

<style scoped>
/* Ajustes finos para el badge circular */
:deep(.v-badge__badge) {
  height: 10px;
  width: 10px;
  min-width: 10px;
}
</style>