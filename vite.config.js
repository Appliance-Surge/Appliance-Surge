import { defineConfig } from 'vite'
import path from 'path'

export default defineConfig({
    root: './',
    resolve: {
        alias: {
            '@': path.resolve(__dirname, './resources'),
        },
    },
    build: {
        outDir: './static',
        rollupOptions: {
            input: {
                main: path.resolve(__dirname, 'resources/js/main.js')
            }
        },
        manifest: true,
    },
})
