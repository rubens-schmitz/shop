/// <reference types="@sveltejs/kit" />

interface PostCategoryRequest {
	title: string;
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

type ListActionFn = (id: number) => Promise<any>;
type ListGetElementsFn = (
    limit: number,
	offset: number,
	title: string,
	categoryId: number
) => Promise<any[]>;

interface ListElement {
	id: number;
	title: string;
	price: number;
	pictures: string[];
	productId?: number;
	quantity?: number;
}

type FormActionFn = (request: any) => void;

type MakeCategoryRequest = (id: number, title: string) => any;

type MakeProductRequest = (
	id: number,
	title: string,
	price: number,
	categoryId: number,
	pictures: string[]
) => any;
