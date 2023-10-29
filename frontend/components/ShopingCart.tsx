import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faTimes } from "@fortawesome/free-solid-svg-icons";
import { cartActions } from "../store/cartActions";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../store";
import { Link } from "react-router-dom";

function ShopingCart() {
  // const [cartItems, setCartItems] = useState([])
  // const [subtotal, setSubtotal] = useState(0)

  // useEffect(() => {
  //   setCartItems(cart)
  //   setSubtotal(getCartSubTotal(cart))
  // }, [cart])
  const dispatch = useDispatch();
  const cartItems = useSelector((state: RootState) => state.cart.itemsList);
  const subTotal = useSelector((state: RootState) => state.cart.totalPrice);

  function updateItem(hendle: string, quantity: number) {
    dispatch(cartActions.updateCartQuantity({ hendle, quantity }));
    // cartActions.updateCartQuantity({ hendle, quantity });
  }

  const removeFromCart = (hendle: string) => {
    dispatch(cartActions.removeCart({ hendle }));
  };

  return (
    <div className="min-h-80 max-w-2xl my-4 sm:my-8 mx-auto w-full">
      <table className="mx-auto">
        <thead>
          <tr className="uppercase text-xs sm:text-sm text-palette-primary border-b border-palette-light">
            <th className="font-primary font-normal px-6 py-4">Product</th>
            <th className="font-primary font-normal px-6 py-4">Quantity</th>
            <th className="font-primary font-normal px-6 py-4 hidden sm:table-cell">
              Price
            </th>
            <th className="font-primary font-normal px-6 py-4">Remove</th>
          </tr>
        </thead>
        <tbody className="divide-y divide-palette-lighter">
          {cartItems.map((item, key) => (
            <tr
              key={key}
              className="text-sm sm:text-base text-gray-600 text-center"
            >
              <td className="font-primary font-medium px-4 sm:px-6 py-4 flex items-center">
                <img
                  src="https://m.media-amazon.com/images/I/61pLuKBfGhL._AC_SL1500_.jpg"
                  alt={item.handle}
                  height={64}
                  width={64}
                  className={`hidden sm:inline-flex`}
                />
                <Link to={`/products/${item.handle}`}>
                  <a className="pt-1 hover:text-palette-dark">{item.title}</a>
                </Link>
              </td>
              <td className="font-primary font-medium px-4 sm:px-6 py-4">
                <input
                  type="number"
                  inputMode="numeric"
                  id="variant-quantity"
                  name="variant-quantity"
                  min="1"
                  step="1"
                  value={item.quantity}
                  onChange={(e) =>
                    updateItem(item.handle, Number(e.target.value))
                  }
                  className="text-gray-900 form-input border border-gray-300 w-16 rounded-sm focus:border-palette-light focus:ring-palette-light"
                />
              </td>
              <td className="font-primary text-base font-light px-4 sm:px-6 py-4 hidden sm:table-cell">
                $<span>{item.price}</span>
              </td>
              <td className="font-primary font-medium px-4 sm:px-6 py-4">
                <button
                  aria-label="delete-item"
                  className=""
                  onClick={() => removeFromCart(item.handle)}
                >
                  <FontAwesomeIcon
                    icon={faTimes}
                    className="w-8 h-8 border border-palette-primary p-1 hover:bg-fuchsia-900"
                  />
                </button>
              </td>
            </tr>
          ))}
          {!cartItems ? null : (
            <tr className="text-center">
              <td></td>
              <td className="font-primary text-base text-gray-600 font-semibold uppercase px-4 sm:px-6 py-4">
                Subtotal
              </td>
              <td className="font-primary text-lg text-palette-primary font-medium px-4 sm:px-6 py-4">
                $<span>{subTotal}</span>
              </td>
              <td></td>
            </tr>
          )}
        </tbody>
      </table>
    </div>
  );
}

export default ShopingCart;
