import React, { useState } from "react";

import './Menu.css'
import Logo from '../Logo/Logo';
import ButOptions, {States} from '../ButOptions/ButOptions';

const Menu = ({format, force}) => {

    const [state, setState] = useState(States.None);

    // global state of application (type of menu)
    function changeOptions(event) {
        setState(event.target.getAttribute('type'));
    }
    
    return (
        <div className="menu">
            <Logo />
            <div className="updButton">
                <button onClick={changeOptions} type={States.Update}>Update Profile</button>
            </div>
            <div className="manageButton">
                <button onClick={changeOptions} type={States.Manage}>Manage Profiles</button>
            </div>
            <div className="readButton">
                <button onClick={changeOptions} type={States.Add}>Find Profile</button>
            </div>
            <div className="deleteButton">
                <button onClick={changeOptions} type={States.Delete}>Delete Profile</button>
            </div>
            <ButOptions state={state} format={format} force={force}/>
        </div>
    )
}
   
export default Menu;