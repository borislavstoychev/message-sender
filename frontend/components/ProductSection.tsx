import React from "react";
import { Product } from "./ProductCard";
import ProductDetails from "./ProductDetails";
import ProductImage, { Image } from "./ProductImage";
import Navigation from "./Navigation";
import Footer from "./Footer";

function ProductSection(productData: Product) {
  const images: Image[] = [
    { src: productData.imageSrc, altText: "test" },
    { src: productData.imageSrc, altText: "test" },
    { src: productData.imageSrc, altText: "test" },
  ];
  return (
    <div className="min-h-screen py-12 sm:pt-20">
      <div className="flex flex-col justify-center items-center md:flex-row md:items-start space-y-8 md:space-y-0 md:space-x-4 lg:space-x-8 max-w-6xl w-11/12 mx-auto">
        <ProductImage images={images} />
        <ProductDetails {...productData} />
      </div>
    </div>
  );
}

export default ProductSection;
