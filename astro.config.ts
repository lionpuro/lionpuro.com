// @ts-check
import { defineConfig, envField } from "astro/config";
import tailwindcss from "@tailwindcss/vite";
import mdx from "@astrojs/mdx";

export default defineConfig({
	site: "https://lionpuro.com",
	trailingSlash: "always",
	devToolbar: {
		enabled: false,
	},
	env: {
		schema: {
			SITE_NAME: envField.string({ context: "server", access: "public" }),
			GITHUB: envField.string({ context: "server", access: "public" }),
			LINKEDIN: envField.string({ context: "server", access: "public" }),
			EMAIL: envField.string({ context: "server", access: "public" }),
		},
	},
	integrations: [mdx()],
	vite: {
		plugins: [tailwindcss()],
	},
	markdown: {
		shikiConfig: {
			themes: {
				light: "github-light",
				dark: "github-dark",
			},
		},
	},
});
