import React from "react";
import BackToProductButton from "../components/BackToProductbutton";
import ShopingCart from "../components/ShopingCart";

function Cart() {
  return (
    <div className="container mx-auto mb-20 min-h-screen">
      {/* <SEO title={pageTitle} />
      <PageTitle text="Your Cart" /> */}
      <ShopingCart />
      <div className="max-w-sm mx-auto space-y-4 px-2">
        {/* <CheckOutButton webUrl={checkoutUrl} /> */}
        <BackToProductButton />
      </div>
    </div>
  );
}

export default Cart;
