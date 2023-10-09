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
    imageSrc: "https://m.media-amazon.com/images/I/61pLuKBfGhL._AC_SL1500_.jpg",
  },
  {
    title: "dog",
    handle: "string",
    description: "Very big Dog!",
    price: 200,
    imageSrc:
      "https://i.etsystatic.com/43546135/r/il/bd6e45/5118617260/il_fullxfull.5118617260_rvua.jpg",
  },
  {
    title: "dog",
    handle: "string",
    description: "Very big Dog!",
    price: 200,
    imageSrc:
      "https://i.etsystatic.com/43546135/r/il/bd6e45/5118617260/il_fullxfull.5118617260_rvua.jpg",
  },
  {
    title: "dog",
    handle: "string",
    description: "Very big Dog!",
    price: 200,
    imageSrc:
      "https://i.etsystatic.com/43546135/r/il/bd6e45/5118617260/il_fullxfull.5118617260_rvua.jpg",
  },
];

export default function Home() {
  return (
    <div className="flex flex-col justify-between min-h-screen">
      <Navigation />
      <Producs data={products} />
      <Footer />
    </div>
  );
}
