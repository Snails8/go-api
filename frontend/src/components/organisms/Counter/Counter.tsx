import React, { useReducer } from "react";

type Action = 'DECREMENT' | 'INCREMENT' | 'DOUBLE' | 'RESET';
const reducer = (currentCount: number, action: Action) => {
    switch (action) {
        case 'INCREMENT':
            return currentCount + 1;
        case 'DECREMENT':
            return currentCount - 1;
        case 'DOUBLE':
            return currentCount * 2
        case 'RESET':
            return 0;
        default:
            return currentCount;
    };
}


type CounterProps = {
    initialProps: number
}

export const Counter = (
    props: CounterProps
) => {

    const { initialProps } = props;
    const [ count, dispatch ] = useReducer(reducer, initialProps)

    return (
        <div>
            <p>{count}</p>
            {/* dispatch にactionを渡して状態を更新 */}
            <button onClick={() => dispatch('DECREMENT')}>-</button>
            <button onClick={() => dispatch('INCREMENT')}>+</button>
            <button onClick={() => dispatch('DOUBLE')}>×2</button>
            <button onClick={() => dispatch('RESET')}>Reset</button>
        </div>
    )
}
