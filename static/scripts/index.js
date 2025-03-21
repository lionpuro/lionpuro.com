document.querySelector("#link-projects")?.addEventListener("click", (e) => {
	e.preventDefault();
	document.querySelector("#projects")?.scrollIntoView({ behavior: "smooth" });
});
