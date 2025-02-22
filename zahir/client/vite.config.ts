import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  build: {
    outDir: './dist',
  },
  server: {
    port: 5174,
    proxy: {
      '/v1': 'http://localhost:8080',
      '/v1/ws': 'ws://localhost:8080',
    }
  }
})
