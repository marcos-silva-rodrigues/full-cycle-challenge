import { Module } from '@nestjs/common';
import { ProductsService } from './products.service';
import { ProductsController } from './products.controller';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { join } from 'path';

@Module({
  imports: [
    ClientsModule.registerAsync([
      {
        name: 'PRODUCT_PACKAGE',
        useFactory: () => ({
          transport: Transport.GRPC,
          options: {
            url: 'grpc-server:50051',
            package: 'github.com.codeedu.codepix',
            protoPath: [join(__dirname, 'proto/product.proto')],
          },
        }),
      },
    ]),
  ],
  controllers: [ProductsController],
  providers: [ProductsService],
})
export class ProductsModule {}
