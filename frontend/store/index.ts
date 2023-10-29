import { configureStore } from "@reduxjs/toolkit";
import cartSlice from "./cartActions";
// import authSlice from "./auth-slice";
// import uiSlice from "./ui-slice";

const store = configureStore({
  reducer: {
    // auth: authSlice.reducer,
    cart: cartSlice.reducer,
    // ui: uiSlice.reducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export default store;
