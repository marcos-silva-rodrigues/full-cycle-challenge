import { Injectable, OnModuleInit, Inject } from '@nestjs/common';
import { CreateProductDto } from './dto/create-product.dto';
import { ProductClientGrpc } from './entities/product.grpc';
import { ClientGrpc } from '@nestjs/microservices';

@Injectable()
export class ProductsService implements OnModuleInit {
  private productGrpcService: ProductClientGrpc;

  constructor(
    @Inject('PRODUCT_PACKAGE')
    private productGrpcPackage: ClientGrpc,
  ) {}

  onModuleInit() {
    this.productGrpcService =
      this.productGrpcPackage.getService('ProductService');
  }

  create(createProductDto: CreateProductDto) {
    return this.productGrpcService.createProduct(createProductDto);
  }

  findAll() {
    return this.productGrpcService.findProducts({});
  }
}
