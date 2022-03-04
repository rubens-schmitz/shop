/**
 * @param {PostCategoryRequest} category
 */
export async function postCategory(category) {
	const body = new FormData();
	body.append('title', category.title);
	return await fetch('/api/category', { method: 'POST', body });
}

/**
 * @param {any} rawCategory
 * @return {GetCategoryResponse}
 */
function makeCategory(rawCategory) {
	return {
		id: parseInt(rawCategory.id),
		title: rawCategory.title
	};
}

/**
 * @param {number} offset
 * @param {string} title
 * @return {Promise<GetCategoryResponse[]>}
 */
export async function getCategories(limit = 0, offset = 0, title = '') {
	let params = `?offset=${offset}&title=${title}&limit=${limit}`;
	const res = await fetch(`/api/categories${params}`);
	const rawCategories = await res.json();
	let categories = [];
	for (let i = 0; i < rawCategories.length; i++)
		categories.push(makeCategory(rawCategories[i]));
	return categories;
}

/**
 * @param {number} id
 * @return {Promise<GetCategoryResponse>}
 */
export async function getCategory(id) {
	const res = await fetch(`/api/category/?id=${id}`);
	const rawCategory = await res.json();
	return makeCategory(rawCategory);
}

/**
 * @param {PutCategoryRequest} category
 */
export async function putCategory(category) {
	const body = new FormData();
	body.append('id', String(category.id));
	body.append('title', category.title);
	return await fetch('/api/category', { method: 'PUT', body });
}

/**
 * @param {number} id
 */
export async function deleteCategory(id) {
	const body = new FormData();
	body.append('id', String(id));
	return await fetch('/api/category', { method: 'DELETE', body });
}
