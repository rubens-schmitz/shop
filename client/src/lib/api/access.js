/**
 * @return {Promise<AdminExistResponse>}
 */
export async function adminExist() {
	const res = await fetch(`/api/access/admin`);
	const raw = await res.json();
	return makeAdminExistResponse(raw);
}

/**
 * @param {any} raw
 */
function makeAdminExistResponse(raw) {
	return {
		success: raw.success
	};
}

/**
 * @return {Promise<CreateAdminResponse>}
 */
export async function createAdmin() {
	const res = await fetch('/api/access/admin', { method: 'POST' });
	const raw = await res.json();
	return makeCreateAdminResponse(raw);
}

/**
 * @param {any} raw
 * @return {CreateAdminResponse}
 */
function makeCreateAdminResponse(raw) {
	return {
		qrcode: raw.qrcode
	};
}

/**
 * @param {string} qrcode
 * @return {Promise<LoginAdminResponse>}
 */
export async function loginAdmin(qrcode) {
	const body = new FormData();
	body.append('qrcode', qrcode);
	const res = await fetch('/api/access/admin', { method: 'PUT', body });
	const raw = await res.json();
	return makeLoginAdminResponse(raw);
}

/**
 * @param {any} raw
 * @return {LoginAdminResponse}
 */
function makeLoginAdminResponse(raw) {
	return {
		success: raw.success
	};
}
