import React from "react";
import { Link } from "react-router-dom";

type Product = {
  title: string;
  handle: string;
  description: string;
  price: number;
  imageSrc: string;
};

function ProductCard(product: Product) {
  const handle = product.handle;
  const title = product.title;
  const description = product.description;
  const price = product.price;

  const imageNode = product.imageSrc;
  console.log(imageNode);
  return (
    <Link
      to={`/products/${handle}`}
      className="h-120 w-72 rounded shadow-lg mx-auto border border-palette-lighter"
    >
      <div className="h-72 border-b-2 border-palette-lighter relative">
        <img
          src={imageNode}
          alt=""
          className="transform duration-500 ease-in-out hover:scale-110"
        />
      </div>
      <div className="h-48 relative">
        <div className="font-primary text-palette-primary text-2xl pt-4 px-4 font-semibold">
          {title}
        </div>
        <div className="text-lg text-gray-600 p-4 font-primary font-light">
          {description}
        </div>
        <div
          className="text-palette-dark font-primary font-medium text-base absolute bottom-0 right-0 mb-4 pl-8 pr-4 pb-1 pt-2 bg-palette-lighter 
            rounded-tl-sm triangle"
        >
          "$"
          <span className="text-lg">{price}</span>
        </div>
      </div>
    </Link>
  );
}

export default ProductCard;
