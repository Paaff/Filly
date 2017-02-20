export class FileItem {
	constructor(
		public name: string,
		public path: string,
		public size: number,
		public type: string
	) {
		console.log(arguments)
	}

	print() {
		console.log(name)
	}
}
