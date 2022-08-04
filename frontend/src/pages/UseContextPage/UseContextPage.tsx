import React, { createContext, useContext, useState } from "react";
import { BasePage } from "../../components/templates/BasePage/BasePage";

export const UserCount = createContext(100);

export const UserUseStateCount = createContext({} as StateCountContext);
type StateCountContext = {
    stateCount: number,
    setStateCount:  React.Dispatch<React.SetStateAction<number>>
}

export const UseContextPage:React.FC =() => {
    // useState併用
    const [stateCount, setStateCount] = useState(1);

    return (
        <>
        <BasePage>
            <h2>useContext検証</h2>
            <p>以下に 親 → 子 → 孫 → ひ孫と表示している</p>
            <UserCount.Provider value={100}>
                <ComponentParent />
            </UserCount.Provider>

            <h3>useStateを使用したパターン</h3>
            {/* useState併用 */}
            <UserUseStateCount.Provider value={{ stateCount, setStateCount }}>
                <ComponentParent />
            </UserUseStateCount.Provider>
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
    const { stateCount, setStateCount } = useContext(UserUseStateCount);

    return (
        <div>
            <h4>ComponentGreatGrandChild</h4>
            <p>親から受け取る値(方法1)：{count}</p>
            <UserCount.Consumer>
                {(count) => {
                    return <p>親から受け取る値(方法2)：{count}</p>;
                }}
            </UserCount.Consumer>

            <p>useState併用パターン：{count}</p>
            <button onClick={() => setStateCount(stateCount + 1)}>+</button>

        </div>
    )
}