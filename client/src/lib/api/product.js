/**
 * @param {PostProductRequest} product
 */
export async function postProduct(product) {
	const body = new FormData();
	body.append('title', product.title);
	body.append('price', String(product.price));
	body.append('categoryId', String(product.categoryId));
	body.append('pictures', String(product.pictures.length));
	for (let i = 0; i < product.pictures.length; i++)
		body.append(`pictures${i}`, product.pictures[i]);
	return await fetch('/api/product', { method: 'POST', body });
}

/**
 * @param {any} rawProduct
 * @return {GetProductResponse}
 */
function makeProduct(rawProduct) {
	return {
		id: parseInt(rawProduct.id),
		title: rawProduct.title,
		price: parseFloat(rawProduct.price),
		categoryId: parseInt(rawProduct.categoryId),
		pictures: rawProduct.pictures
	};
}

/**
 * @param {number} offset
 * @param {string} title
 * @param {number} category
 * @return {Promise<GetProductResponse[]>}
 */
export async function getProducts(
	limit = 0,
	offset = 0,
	title = '',
	category = 0
) {
	let params = `?limit=${limit}&offset=${offset}&title=${title}`;
    params += `&categoryId=${category}`
	const res = await fetch(`/api/products/${params}`);
	const rawProducts = await res.json();
	let products = [];
	for (let i = 0; i < rawProducts.length; i++)
		products.push(makeProduct(rawProducts[i]));
	return products;
}

/**
 * @param {number} id
 * @return {Promise<GetProductResponse>}
 */
export async function getProduct(id) {
	const res = await fetch(`/api/product/?id=${id}`);
	const rawProduct = await res.json();
	return makeProduct(rawProduct);
}

/**
 * @param {PutProductRequest} product
 */
export async function putProduct(product) {
	const body = new FormData();
	body.append('id', String(product.id));
	body.append('title', product.title);
	body.append('price', String(product.price));
	body.append('categoryId', String(product.categoryId));
	body.append('pictures', String(product.pictures.length));
	for (let i = 0; i < product.pictures.length; i++)
		body.append(`pictures${i}`, product.pictures[i]);
	return await fetch('/api/product', { method: 'PUT', body });
}

/**
 * @param {number} id
 */
export async function deleteProduct(id) {
	const body = new FormData();
	body.append('id', String(id));
	return await fetch('/api/product', { method: 'DELETE', body });
}
