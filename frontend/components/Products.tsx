import React from "react";
import ProductCard from "./ProductCard";

export type Product = {
  title: string;
  handle: string;
  description: string;
  price: number;
  imageSrc: string;
};

type ProductProps = {
  data: Product[];
};

function Producs({ data }: ProductProps) {
  return (
    <div className="py-12 max-w-6xl mx-auto grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-x-4 gap-y-8">
      {data.map((product, index) => (
        <ProductCard key={index} {...product} />
      ))}
    </div>
  );
}

export default Producs;
