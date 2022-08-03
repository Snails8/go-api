import React from "react";
import { useSearchParams } from "react-router-dom";
import { BasePage } from "../../components/templates/BasePage/BasePage";

export const UseContextPage:React.FC =() => {

    return (
        <>
        <BasePage>
            <h3>useContext検証</h3>

            <button onClick={() => window.open("page3")}>別タブで開く</button>
        </BasePage>
        </>
    )
}