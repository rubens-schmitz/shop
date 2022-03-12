/**
 * @param {any} rawCart
 * @return {GetCartResponse}
 */
export function makeCart(rawCart) {
	return {
		price: parseFloat(rawCart.price),
		quantity: parseInt(rawCart.quantity)
	};
}

/**
 * @return {Promise<GetCartResponse>}
 */
export async function getCart() {
	const res = await fetch(`/api/cart`);
	const rawDeal = await res.json();
	return makeCart(rawDeal);
}
