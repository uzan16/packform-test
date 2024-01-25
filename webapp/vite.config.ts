import { fileURLToPath, URL } from 'node:url'

import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
  const {
    VITE_PORT,
    VITE_BASE_API_URL,
  } = loadEnv(mode, process.cwd());

  return {
    plugins: [
      vue(),
      vueJsx(),
    ],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url))
      }
    },
    server: {
      open: true,
      port: parseInt(VITE_PORT),
      proxy: {
        '/api': {
          target: VITE_BASE_API_URL,
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/api/, ''),
        },
      },
    },
  };
});