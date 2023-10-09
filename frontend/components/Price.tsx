function Price(currency: string, num: number, numSize: string) {
  return (
    <>
      {currency}
      <span className={numSize}>{num}</span>
    </>
  );
}

export default Price;
