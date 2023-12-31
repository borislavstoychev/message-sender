import React from "react";
import { useState, useRef } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faArrowLeft, faArrowRight } from "@fortawesome/free-solid-svg-icons";

export type Image = {
  src: string;
  altText: string;
};

type ProductImages = {
  images: Image[];
};

function ProductImage({ images }: ProductImages) {
  const [mainImg, setMainImg] = useState(images[0]);
  const ref = useRef<HTMLDivElement>(null);

  function scroll(scrollOffset: number) {
    ref.current!.scrollLeft += scrollOffset;
  }

  return (
    <div className="w-full md:w-1/2 max-w-md border border-palette-lighter bg-white rounded shadow-lg">
      <div className="relative h-96">
        <img
          src={mainImg.src}
          alt={mainImg.altText}
          className="transform duration-500 ease-in-out hover:scale-105"
        />
      </div>
      <div className="relative flex border-t border-palette-lighter">
        <button
          aria-label="left-scroll"
          className="h-32 bg-palette-lighter hover:bg-palette-light  absolute left-0 z-10 opacity-75"
          onClick={() => scroll(-300)}
        >
          <FontAwesomeIcon
            icon={faArrowLeft}
            className="w-3 mx-1 text-palette-primary"
          />
        </button>
        <div
          ref={ref}
          style={{ scrollBehavior: "smooth" }}
          className="flex space-x-1 w-full overflow-auto border-t border-palette-lighter"
        >
          {images.map((imgItem, index) => (
            <button
              key={index}
              className="relative w-40 h-32 flex-shrink-0 rounded-sm "
              onClick={() => setMainImg(images[index])}
            >
              <img
                src={imgItem.src}
                alt={imgItem.altText}
                className="object-fill"
              />
            </button>
          ))}
        </div>
        <button
          aria-label="right-scroll"
          className="h-32 bg-palette-lighter hover:bg-palette-light  absolute right-0 z-10 opacity-75"
          onClick={() => scroll(300)}
        >
          <FontAwesomeIcon
            icon={faArrowRight}
            className="w-3 mx-1 text-palette-primary"
          />
        </button>
      </div>
    </div>
  );
}

export default ProductImage;
