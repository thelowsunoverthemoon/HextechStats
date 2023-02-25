import React, { useEffect, useState, useRef } from "react";

import Button from 'react-bootstrap/Button';

import ProfList from '../ProfList/ProfList';
import ProfChoose from '../ProfChoose/ProfChoose';

import { behavDelete, behavUpdate } from "../Operations/Operations.js";

import './ButOptions.css';

export const States = {
    None: 0,
	Update: 1,
	Add: 2,
	Manage: 3,
	Delete: 4
}

const ButOptions = ({state, format, force}) => {
    const added = useRef([]);
    const updated = useRef([]);
    const deleted = useRef([]);

    const [option, setOption] = useState(<div></div>);

    useEffect(() => {

        // reset these everytime we reload element
        updated.current = [];
        deleted.current = [];

        if (state == States.Add) {
            setOption(<ProfChoose format={format} force={force} added={added}/>);
        } else if (state == States.Update) {
            setOption(
                <div className="listContainer">
                    <Button onClick={() => behavUpdate(updated.current, force, format)}>Update</Button>
                    <ProfList key="1" type={state} clkArr={updated} />
                </div>
            );
        } else if (state == States.Delete) {
            setOption(
                <div className="listContainer">
                    <Button onClick={() => behavDelete(deleted.current, force, format)}>Delete</Button>
                    <ProfList key="2" type={state} clkArr={deleted} />
                </div>
            );
        } else {
            setOption(<ProfList key="3" type={state} clkArr={added} format={format} force={force} />);
        }
    }, [state]);

    return (
        <>
            {option}
        </>
    )
}
   
export default ButOptions;