import { dialog } from './stores';

/**
 * @type {string}
 */
export const modalOkButton = 'modalOkButton';

/**
 * @type {Dialog}
 */
export const emptyDialog = {
	task: '',
	body: '',
	qrcode: '',
	reload: false
};

/**
 * @param {Dialog} props
 */
export function openModal(props) {
	let o = emptyDialog;
	Object.assign(o, props);
	dialog.set(o);
	setTimeout(() => {
		document.getElementById(modalOkButton).focus();
	}, 100);
}
