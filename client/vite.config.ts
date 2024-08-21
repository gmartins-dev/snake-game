import react from '@vitejs/plugin-react';
import dotenv from 'dotenv';
import { defineConfig } from 'vite';
import envCompatible from 'vite-plugin-env-compatible';

dotenv.config();

export default defineConfig({
  plugins: [react(), envCompatible()],
  server: {
    proxy: {
      '/api': {
        target: `http://${process.env.SERVER_ADDRESS}`,
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ''),
      },
    },
  },
});
