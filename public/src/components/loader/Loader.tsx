import React from "react";
import loader from "./Loader.module.css"

export default function Loader() {
  return (
    <div className={loader["lds-hourglass"]}>
    </div>
  );
}