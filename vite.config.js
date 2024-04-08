import { defineConfig } from "vite";
import { resolve } from "path";

export default defineConfig({
    build: {
        outDir: "build",
        rollupOptions: {
            input: {
                main: resolve(__dirname, "cmd/bbs/main.go"),
            },
        },
    },
    appType: "custom",
    css: {},
    assetsInclude: ["**/*.templ"],
});
