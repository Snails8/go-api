import React, { createContext, useContext } from "react";
import { BasePage } from "../../components/templates/BasePage/BasePage";

export const UserCount = createContext(100);

export const UseContextPage:React.FC =() => {

    return (
        <>
        <BasePage>
            <h2>useContext検証</h2>
            <p>以下に 親 → 子 → 孫 → ひ孫と表示している</p>
            <UserCount.Provider value={100}>
                <ComponentParent />
            </UserCount.Provider>
        </BasePage>
        </>
    )
}

// 親
export const ComponentParent:React.FC = ({

}) => {
    return (
        <div>
            <h4>Component Parent</h4>
            <ComponentChild />
        </div>
    )
}

// 子
export const ComponentChild:React.FC = ({

}) => {
    return (
        <div>
            <h4>Component Child</h4>
            <ComponentGrandChild />
        </div>
    )
}

// 孫
export const ComponentGrandChild:React.FC = ({

}) => {
    return (
        <div>
            <h4>ComponentGrandChild</h4>
            <ComponentGreatGrandChild />
        </div>
    )
}

// ひ孫
export const ComponentGreatGrandChild:React.FC = ({

}) => {
    const count = useContext(UserCount);

    return (
        <div>
            <h4>ComponentGreatGrandChild</h4>
            <p>親から受け取る値(方法1)：{count}</p>
            <UserCount.Consumer>
                {(count) => {
                    return <p>親から受け取る値(方法2)：{count}</p>;
                }}
            </UserCount.Consumer>
        </div>
    )
}