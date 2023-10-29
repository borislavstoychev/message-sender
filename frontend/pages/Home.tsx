import React from "react";
import Navigation from "../components/Navigation";
import Footer from "../components/Footer";
import Producs, { Product } from "../components/Products";

const products: Product[] = [
  {
    title: "dog",
    handle: "bobby",
    description: "Very big Dog!",
    price: 200,
    imageSrc: "https://m.media-amazon.com/images/I/61pLuKBfGhL._AC_SL1500_.jpg",
  },
  {
    title: "dog",
    handle: "sonq",
    description: "Very big Dog!",
    price: 200,
    imageSrc:
      "https://images.unsplash.com/photo-1583336663277-620dc1996580?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mnx8ZG9nJTIwY2xvdGhlc3xlbnwwfHwwfHx8MA%3D%3D&w=1000&q=80",
  },
  {
    title: "ivo",
    handle: "ivo",
    description: "Very big Dog!",
    price: 200,
    imageSrc:
      "https://i.etsystatic.com/6144628/r/il/836a55/2008597499/il_fullxfull.2008597499_5fzq.jpg",
  },
  {
    title: "dog",
    handle: "test",
    description: "Very big Dog!",
    price: 200,
    imageSrc:
      "https://i.etsystatic.com/43546135/r/il/bd6e45/5118617260/il_fullxfull.5118617260_rvua.jpg",
  },
];

export default function Home() {
  return (
    <div className="flex flex-col justify-between min-h-screen">
      <Producs data={products} />
    </div>
  );
}
