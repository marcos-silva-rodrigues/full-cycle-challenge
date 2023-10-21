import { Observable } from 'rxjs';

interface Product {
  id: number;
  name: string;
  description: string;
  price: number;
}

export interface ProductListRpcResponse {
  products: Product[];
}

export interface ProductClientGrpc {
  createProduct: (data: {
    name: string;
    description: string;
    price: number;
  }) => Observable<{ id: string; status: string; error: string }>;
  findProducts: () => Observable<ProductListRpcResponse>;
}
