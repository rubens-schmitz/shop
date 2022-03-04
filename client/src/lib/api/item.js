/**
 * @param {number} productId
 */
export async function postItem(productId) {
	const body = new FormData();
	body.append('productId', String(productId));
	return await fetch('/api/item', { method: 'POST', body });
}

/**
 * @param {any} rawItem
 * @return {GetItemResponse}
 */
function makeItem(rawItem) {
	return {
		id: parseInt(rawItem.id),
		title: rawItem.title,
		price: parseFloat(rawItem.price),
		categoryId: parseInt(rawItem.categoryId),
		pictures: rawItem.pictures,
		productId: parseInt(rawItem.productId),
		quantity: parseInt(rawItem.quantity)
	};
}

/**
 * @return {Promise<GetItemResponse[]>}
 */
export async function getItems(
	limit = 0,
	offset = 0,
	title = '',
	categoryId = 0
) {
	let params = `?limit=${limit}&offset=${offset}&title=${title}&categoryId=${categoryId}`;
	const res = await fetch('/api/items' + params);
	const rawItems = await res.json();
	let items = [];
	for (let i = 0; i < rawItems.length; i++)
		items.push(makeItem(rawItems[i]));
	return items;
}

/**
 * @param {number} id
 * @param {number} quantity
 */
export async function putItem(id, quantity) {
	const body = new FormData();
	body.append('id', String(id));
	body.append('quantity', String(quantity));
	return await fetch('/api/item', { method: 'PUT', body });
}

/**
 * @param {number} id
 */
export async function deleteItem(id) {
	const body = new FormData();
	body.append('id', String(id));
	return await fetch('/api/item', { method: 'DELETE', body });
}
