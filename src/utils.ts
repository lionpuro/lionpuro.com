import { getCollection } from "astro:content";

export async function fetchPosts() {
	const posts = (await getCollection("posts")).sort(
		(a, b) => b.data.date.valueOf() - a.data.date.valueOf(),
	);
	return posts;
}

export function dateString(d: Date): string {
	return d.toLocaleDateString("default", {
		month: "long",
		day: "numeric",
		year: "numeric",
	});
}
