import { join, resolve } from "path";
import { defineConfig } from "vite";
import { viteStaticCopy } from "vite-plugin-static-copy";

export default defineConfig({
	build: {
		emptyOutDir: true,
		outDir: join(__dirname, "assets/public/scripts"),
		rollupOptions: {
			external: [],
			input: [
				resolve(__dirname, "assets/src/modal-image.ts"),
			],
			output: {
				entryFileNames: "[name].js",
			},
		},
	},
	plugins: [
		viteStaticCopy({
			targets: [
				{
					src: resolve(__dirname, "assets/static/*"),
					dest: resolve(__dirname, "assets/public"),
				},
			],
		}),
	],
});
