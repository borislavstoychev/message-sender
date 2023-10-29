// import { useState } from 'react'
// import ProductInfo from '@/components/ProductInfo'
import React from "react";
import ProductForm from "./ProdictForm";
import { Product } from "./ProductCard";
import BackToProductButton from "./BackToProductbutton";

function ProductDetails(product: Product) {
  // const [variantPrice, setVariantPrice] = useState(productData.variants.edges[0].node.price)

  return (
    <div className="flex flex-col justify-between h-full w-full md:w-1/2 max-w-xs mx-auto space-y-4 min-h-128">
      <BackToProductButton />
      <div className="h-48 relative">
        <div className="font-primary text-palette-primary text-2xl pt-4 px-4 font-semibold">
          {product.title}
        </div>
        <div className="text-lg text-gray-600 p-4 font-primary font-light">
          {product.description}
        </div>
        <div
          className="text-purple-900 font-primary font-medium text-xl absolute bottom-0 right-0 mb-4 pl-8 pr-4 pb-1 pt-2 bg-palette-lighter 
            rounded-tl-sm triangle"
        >
          $<span>{product.price}</span>
        </div>
      </div>
      <ProductForm {...product} />
    </div>
  );
}

export default ProductDetails;
