import React, { useReducer } from "react";
import { Counter } from "../../components/organisms/Counter/Counter";
import { BasePage } from "../../components/templates/BasePage/BasePage";

export const UseReducerPage:React.FC =() => {

    return (
        <>
        <BasePage>
            <h2>useReducer検証</h2>
            <Counter
                initialProps={1}
            ></Counter>
        </BasePage>
        </>
    )
}