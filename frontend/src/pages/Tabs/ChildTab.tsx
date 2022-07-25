import React from "react";
import { useSearchParams } from "react-router-dom";
import { BasePage } from "../../components/templates/BasePage/BasePage";

export const ChildTabPage:React.FC =() => {

    const [searchParams] = useSearchParams()

    const query1 = searchParams.get("name")
    const query2 = searchParams.get("type");

    return (
        <>
        <BasePage>
            <h3>子タブ</h3>
            <p>name={query1}</p>
            <p>type={query2}</p>
            <button onClick={() => window.open("page3")}>別タブで開く</button>
        </BasePage>
        </>
    )
}