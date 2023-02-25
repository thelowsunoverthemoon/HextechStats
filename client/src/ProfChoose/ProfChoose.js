import React, { useState } from "react";
import Dropdown from 'react-bootstrap/Dropdown';
import DropdownButton from 'react-bootstrap/DropdownButton';
import { behavAdd } from "../Operations/Operations.js";

import './ProfChoose.css'

import 'bootstrap/dist/css/bootstrap.css';

const servers = [
    "North America", "Europe", "Japan", "Korea", "Latin America", "Oceania", "Russia"
];

const ProfChoose = ({format, force, added}) => {
    const [server, setServer] = useState('');
    const [name, setName] = useState('');

    function handleServer(e) {
        setServer(e);
    }

    function makeDropDown(text, def, options, set) {
        return (
            <DropdownButton id="dropdown" title={text === '' ? def : text} onSelect={set}>
                <>
                    {options.map(
                        (op) => (
                            <Dropdown.Item key={op} eventKey={op}>{op}</Dropdown.Item>
                        )
                    )}
                </>
            </DropdownButton>
        );
    }

    return (
        <div className="chooseContainer">
            <p>Enter Info</p>
            {makeDropDown(server, "Server", servers, handleServer)}
            <input
                className="input"
                type="text"
                placeholder="Name"
                value={name}
                onChange={(event) => setName(event.target.value)}
                onKeyPress={(event) => {if (event.key === 'Enter' && server != '') {
                    behavAdd(server, name, format, force, added);
                }}}
            />
        </div>
    )
}

export default ProfChoose;