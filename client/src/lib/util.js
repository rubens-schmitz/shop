/**
 * @return {number}
 */
export function getURLIdParam() {
	const href = window.location.href;
	const re = /.+\?id=([0-9]+)/;
	const match = href.match(re);
	if (match === null) return undefined;
	return parseInt(match[1]);
}

/**
 * @param {string} name
 * @return {string}
 */
export function getCookieValue(name) {
	return (
		document.cookie
			.match('(^|;)\\s*' + name + '\\s*=\\s*([^;]+)')
			?.pop() || ''
	);
}
