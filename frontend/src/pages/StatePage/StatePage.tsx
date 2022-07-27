import React, { useEffect, useState } from "react";
import { BasePage } from "../../components/templates/BasePage/BasePage";

export const StatePage:React.FC =() => {
    const [reloadText, setReloadText] = useState<string>("") 
    
    useEffect(() => {
        setReloadText(localStorage.getItem('reload') as string)
        
        window.addEventListener('beforeunload', (e) => {
            console.log("reloadりろーど時に実行される")
            localStorage.setItem('reload', "reloadされたよ")
        })
    },[])
    return (
        <>
        <BasePage>
            <h3>状態保持の検証</h3>

            <h3>ブラウザのリロードを感知する</h3>
            <span>{reloadText}</span>
        </BasePage>
        </>
    )
}