class ModalImage extends HTMLElement {
	public minWidth?: number;

	constructor() {
		super();
	}

	connectedCallback() {
		const minWidth = this.getAttribute("minWidth");
		if (minWidth && !isNaN(Number(minWidth))) {
			this.minWidth = Number(minWidth);
		}
		const img = this.querySelector("img");
		if (!img) {
			return console.warn("no img found in <modal-image>");
		}
		const dialog = document.createElement("dialog");
		dialog.classList.add("modal-image__dialog");

		const container = document.createElement("div");
		container.classList.add("modal-image__dialog__content");

		const image = document.createElement("img");
		image.classList.add("modal-image__image");

		const close = document.createElement("button");
		close.innerHTML = "×";
		close.classList.add("modal-image__close");
		close.addEventListener("click", () => dialog.close());

		container.appendChild(image);
		dialog.appendChild(container);
		dialog.appendChild(close);

		this.appendChild(dialog);

		img.addEventListener("click", (e) => {
			e.preventDefault();
			if (this.minWidth && window.innerWidth <= this.minWidth) {
				return;
			}
			image.src = img.src;
			dialog.showModal();
		});

		dialog.addEventListener("click", (e) => {
			const dialogRect = dialog.getBoundingClientRect();
			if (
				e.clientX < dialogRect.left ||
				e.clientX > dialogRect.right ||
				e.clientY < dialogRect.top ||
				e.clientY > dialogRect.bottom
			) {
				dialog.close();
			}
		});
	}
}

if (!customElements.get("modal-image")) {
	customElements.define("modal-image", ModalImage);
}
