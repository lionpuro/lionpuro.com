// @ts-check
import { defineConfig } from "astro/config";
import tailwindcss from "@tailwindcss/vite";
import mdx from "@astrojs/mdx";

export default defineConfig({
	site: "https://lionpuro.com",
	trailingSlash: "always",
	devToolbar: {
		enabled: false,
	},
	integrations: [mdx()],
	vite: {
		plugins: [tailwindcss()],
	},
	markdown: {
		shikiConfig: {
			theme: "github-dark",
		},
	},
});
