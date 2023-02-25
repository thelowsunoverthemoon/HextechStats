import React, {useState, useRef, useReducer} from "react";

import Charts from './Charts/Charts';
import Menu from './Menu/Menu';

const App = () => {

    const format = useRef([
        {
            lines : [],
            data: []
        },
        {
            lines : [],
            data: []
        },
        {
            lines : [],
            data: []
        },
        {
            lines : [],
            data: []
        },
        {
            lines : [],
            data: []
        },
        {
            lines : [],
            data: []
        }
    ]);
    // used to update graphs
    const [_, forceUpdate] = useReducer((x) => x + 1, 0);

    return (
        <div>
            <Charts format={format}/>
            <Menu format={format} force={forceUpdate}/>
        </div>
    )
}

export default App;
