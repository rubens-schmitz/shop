/**
 * @return {Promise<PostDealResponse>}
 */
export async function postDeal() {
	const res = await fetch('/api/deal', { method: 'POST' });
	return await res.json();
}
