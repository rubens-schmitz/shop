/**
 * @return {Promise<PostDealResponse>}
 */
export async function postDeal() {
	const res = await fetch('/api/deal', { method: 'POST' });
	return await res.json();
}

/**
 * @param {number} limit
 * @param {number} offset
 * @param {string} title
 * @param {number} categoryId
 * @return {Promise<GetDealResponse[]>}
 */
export async function getDeals(
	limit = 0,
	offset = 0,
	title = '',
	categoryId = 0
) {
	let params = `?limit=${limit}&offset=${offset}&title=${title}`;
	params += `&categoryId=${categoryId}`;
	const res = await fetch('/api/deals' + params);
	const rawDeals = await res.json();
	let deals = [];
	for (let i = 0; i < rawDeals.length; i++)
		deals.push(makeDeal(rawDeals[i]));
	return deals;
}

/**
 * @param {any} rawDeal
 * @return {GetDealResponse}
 */
export function makeDeal(rawDeal) {
	return {
		id: parseInt(rawDeal.id),
		code: String(rawDeal.code),
		datestamp: String(rawDeal.datestamp),
		price: parseFloat(rawDeal.price),
		quantity: parseInt(rawDeal.quantity),
		cartId: parseInt(rawDeal.cartId)
	};
}

/**
 * @param {number} id
 * @return {Promise<GetDealResponse>}
 */
export async function getDeal(id) {
	const res = await fetch(`/api/deal/?id=${id}`);
	const rawDeal = await res.json();
	return makeDeal(rawDeal);
}

/**
 * @param {number} id
 * @return {Promise<any>}
 */
export async function deleteDeal(id) {
	const body = new FormData();
	body.append('id', String(id));
	return await fetch(`/api/deal`, { method: 'DELETE', body });
}
