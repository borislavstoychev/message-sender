import { createSlice, PayloadAction } from "@reduxjs/toolkit";
// import { uiActions } from "./ui-slice";

export type ProductsCard = {
  handle: string;
  price: number;
  quantity: number;
  totalPrice: number;
  title: string;
};

const cartSlice = createSlice({
  name: "cart",
  initialState: {
    itemsList: [] as ProductsCard[],
    totalQuantity: 0,
    totalPrice: 0,
    showCart: false,
    changed: false,
  },
  reducers: {
    replaceData(state, action) {
      state.totalQuantity = action.payload.totalQuantity;
      state.itemsList = action.payload.itemsList;
    },
    addToCart(state, action: PayloadAction<ProductsCard>) {
      state.changed = true;
      const newItem: ProductsCard = action.payload;
      // to check if item is already available
      const existingItem = state.itemsList.find(
        (item: ProductsCard) => item.handle === newItem.handle
      );

      if (existingItem) {
        existingItem.quantity += newItem.quantity;
        existingItem.totalPrice += newItem.price;
        state.totalQuantity += newItem.quantity;
        state.totalPrice += newItem.price;
      } else {
        state.itemsList.push({
          handle: newItem.handle,
          price: newItem.price,
          quantity: newItem.quantity,
          totalPrice: newItem.price * newItem.quantity,
          title: newItem.title,
        });
        state.totalQuantity += newItem.quantity;
        state.totalPrice += newItem.price * newItem.quantity;
      }
      // console.log(state.itemsList);
    },
    updateCartQuantity(
      state,
      action: PayloadAction<{ hendle: string; quantity: number }>
    ) {
      state.changed = true;
      const id = action.payload.hendle;
      const quantity = action.payload.quantity;

      const existingItem = state.itemsList.find((item) => item.handle === id);
      if (existingItem) {
        // state.itemsList = state.itemsList.filter((item) => item.handle !== id);
        existingItem.quantity = quantity;
        existingItem.totalPrice = quantity * existingItem.price;
        state.totalQuantity = state.itemsList.reduce(
          (a, b) => a + b.quantity,
          0
        );
        state.totalPrice = state.itemsList.reduce(
          (a, b) => a + b.totalPrice,
          0
        );
      }
      // else {
      //   existingItem!.quantity -= quantity;
      //   existingItem!.totalPrice -= existingItem!.price * quantity;
      //   state.totalQuantity--;
      //   state.totalPrice -= existingItem!.totalPrice;
      // }
    },
    removeCart(state, action: PayloadAction<{ hendle: string }>) {
      state.changed = true;
      const id = action.payload.hendle;

      const existingItem = state.itemsList.find((item) => item.handle === id);
      if (existingItem) {
        state.itemsList = state.itemsList.filter((item) => item.handle !== id);
        state.totalQuantity -= existingItem.quantity;
        state.totalPrice -= existingItem.totalPrice;
      }
    },
    setShowCart(state) {
      state.showCart = !state.showCart;
    },
  },
});

export const cartActions = cartSlice.actions;

export default cartSlice;
