import React from "react";
import Navigation from "../components/Navigation";
import Footer from "../components/Footer";
import Producs, { Product } from "../components/Products";

const products: Product[] = [
  {
    title: "dog",
    handle: "string",
    description: "Very big Dog!",
    price: 200,
    imageSrc:
      "https://cdn.shopify.com/s/files/1/2800/2014/products/mockup-fc750eaa.jpg?v=1616988549",
  },
  {
    title: "dog",
    handle: "string",
    description: "Very big Dog!",
    price: 200,
    imageSrc:
      "https://cdn.shopify.com/s/files/1/2800/2014/products/mockup-fc750eaa.jpg?v=1616988549",
  },
  {
    title: "dog",
    handle: "string",
    description: "Very big Dog!",
    price: 200,
    imageSrc:
      "https://cdn.shopify.com/s/files/1/2800/2014/products/mockup-fc750eaa.jpg?v=1616988549",
  },
  {
    title: "dog",
    handle: "string",
    description: "Very big Dog!",
    price: 200,
    imageSrc:
      "https://cdn.shopify.com/s/files/1/2800/2014/products/mockup-fc750eaa.jpg?v=1616988549",
  },
];

export default function Home() {
  products.map((product, index) => console.log(product));
  return (
    <div className="flex flex-col justify-between min-h-screen">
      <Navigation />
      <Producs data={products} />
      <Footer />
    </div>
  );
}
