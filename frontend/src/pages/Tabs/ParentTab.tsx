import React, { useEffect, useState } from "react";
import { BasePage } from "../../components/templates/BasePage/BasePage";

export const ParentTabPage:React.FC =() => {
    const [message, setMessage] = useState('message')
    const [childString, setChildString] = useState('')

    useEffect(() => {
        localStorage.setItem('key', JSON.stringify(message));
    },[message])

    const handleChange = (e: any) => {
        setMessage(e.currentTarget.value)
        localStorage.setItem('key', JSON.stringify(message));
    }
    
    useEffect(() => {
        console.log(childString)
    }, [childString])
    return (
        <>
        <BasePage>
            <h3>親タブ</h3>
            <button onClick={() => window.open("child-tab")}>別タブで開く</button>

            <h4>localStorageを使用した親 → 子 の受け渡し方法</h4> <br/>
            <label>text</label>
            <input type="text" name="text" defaultValue={message}
                onChange={handleChange}
            />

            <h4>window.opener を使用して子タブから値を受け取る</h4> <br/>
            <label>text</label>
            <input type="text" name="childString" id="childString" defaultValue={childString}
            />
        </BasePage>    
        </>
    )
}