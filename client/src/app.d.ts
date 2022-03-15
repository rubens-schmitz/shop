/// <reference types="@sveltejs/kit" />

interface PostDealResponse {
	qrcode: string;
}

interface PostCategoryRequest {
	title: string;
}

interface AdminExistResponse {
	success: boolean;
}

interface CreateAdminResponse {
	qrcode: string;
}

interface LoginAdminResponse {
	success: boolean;
}

interface GetCategoryResponse {
	id: number;
	title: string;
}

interface PutCategoryRequest extends PostCategoryRequest {
	id: number;
}

interface PostProductRequest {
	title: string;
	price: number;
	categoryId: number;
	pictures: string[];
}

interface PutProductRequest extends PostProductRequest {
	id: number;
}

interface GetProductResponse {
	id: number;
	title: string;
	price: number;
	categoryId: number;
	pictures: string[];
}

interface GetItemResponse extends GetProductResponse {
	productId: number;
	quantity: number;
}

interface GetDealResponse {
	id: number;
	code: string;
	datestamp: string;
	price: number;
	quantity: number;
	cartId: number;
}

interface GetCartResponse {
	price: number;
	quantity: number;
}

type ListActionFn = (id: number) => Promise<any>;

type ListGetElementsFn = (
	limit: number,
	offset: number,
	title: string,
	categoryId: number
) => Promise<any[]>;

interface ListElement {
	id?: number;
	title?: string;
	price?: number;
	pictures?: string[];
	productId?: number;
	quantity?: number;
	datestamp?: string;
}

type FormActionFn = (request: any) => Promise<any>;

type MakeCategoryRequest = (id: number, title: string) => any;

type MakeProductRequest = (
	id: number,
	title: string,
	price: number,
	categoryId: number,
	pictures: string[]
) => any;
