import React, { useEffect, useRef, useState } from "react";
import { BasePage } from "../../components/templates/BasePage/BasePage";

export const StatePage:React.FC = ({

}) => {
    const [reloadText, setReloadText] = useState<string>("") 
    
    useEffect(() => {
        window.addEventListener('beforeunload', (e) => {
            // console.log("reloadりろーど時に実行される")
            localStorage.setItem('reload', "reloaded")
        })
        
        if (localStorage.getItem('reload') === "reloaded") {
            setReloadText(localStorage.getItem('reload') as string)
        }

        localStorage.setItem('reload', "")
    },[])

    // ref検証
    const [count, setCount] = useState(0);
    const countRef = useRef(0);

    // useEffect 検証 (初期render -> アンマウント -> 第二引数指定)
    useEffect(() => {
        console.log("第二引数[]: 初回時に発火する");
    },[])

    useEffect(() => {
        console.log("第二引数指定: 初回・値変更時に発火する");
    },[count])

    useEffect(() => {
        return () => {
            console.log("クリーンアップ: 値変更時・アンマウント時に発火する");
        }
    },[count])

    useEffect(() => {
        return () => {
            console.log("クリーンアップ: アンマウント時に発火する");
        }
    },[])



    return (
        <>
        <BasePage>
            <h3>状態保持の検証</h3>

            <h3>ブラウザのリロードを感知してリロード時のみ状態を保持する</h3>
            <span>status:{reloadText}</span>

            <h3>遷移元のURLを取得する</h3>
            <span>{window.document.referrer}</span>

            <h3>useRefとuseStateを検証する</h3>
            <div>
                <div>カウント（useState）: {count}</div>
                <button onClick={() => setCount(count + 1)}>カウントアップ（useState）</button>

                <div>カウント（useRef）: {countRef.current}</div>
                <button onClick={() => countRef.current++}>カウントアップ（useRef）</button>
            </div>
        </BasePage>
        </>
    )
}