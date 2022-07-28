import React, { useEffect, useState } from "react";
import { BasePage } from "../../components/templates/BasePage/BasePage";

export const StatePage:React.FC = ({

}) => {
    const [reloadText, setReloadText] = useState<string>("") 
    
    useEffect(() => {
        setReloadText(localStorage.getItem('reload') as string)
        
        window.addEventListener('beforeunload', (e) => {
            console.log("reloadりろーど時に実行される")
            localStorage.setItem('reload', "reloadされたよ")

            //ここでstateを入れてもリセットされる
        })
    },[])

    return (
        <>
        <BasePage>
            <h3>状態保持の検証</h3>

            <h3>ブラウザのリロードを感知する</h3>
            <span>{reloadText}</span>

            <h3>遷移元のURLを取得する</h3>
            <span>{window.document.referrer}</span>
        </BasePage>
        </>
    )
}