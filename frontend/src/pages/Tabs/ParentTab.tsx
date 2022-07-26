import React, { useEffect, useState } from "react";
import { BasePage } from "../../components/templates/BasePage/BasePage";

export const ParentTabPage:React.FC =() => {
    const [message, setMessage] = useState('message')

    useEffect(() => {
        localStorage.setItem('key', JSON.stringify(message));
    },[message])

    const handleChange = (e: any) => {
        setMessage(e.currentTarget.value)
        localStorage.setItem('key', JSON.stringify(message));
    }


    return (
        <>
        <BasePage>
            <h3>親タブ</h3>
            <button onClick={() => window.open("child-tab")}>別タブで開く</button>

            <br/>
            <label>text</label>
            <input type="text" name="text" defaultValue={message}
                onChange={handleChange}
            />
        </BasePage>    
        </>
    )
}