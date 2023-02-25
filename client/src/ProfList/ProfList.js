import React, { useEffect, useState } from "react";
import { States } from "../ButOptions/ButOptions";
import { behavManage, behavEntries } from "../Operations/Operations.js";
import './ProfList.css';

const ProfList = ({type, format, clkArr, force}) => {

    const [entries, setEntries] = useState([]);

    // reload entries everytime it element reloaded
    useEffect(() => {

        behavEntries(setEntries);

    }, []);
    
    function selectRow(e, force, format) {
        e = e || window.event;
        let target = e.srcElement || e.target;
        while (target && target.nodeName !== "TR") {
            target = target.parentNode;
        }

        if (target) {
            let cells = target.getElementsByTagName("td");

            let elem = {name : cells[0].innerHTML, server : cells[1].innerHTML};

            let add = true;
            if (clkArr.current.some(e => e.name == elem.name && e.server == elem.server)) {
                target.style.backgroundColor = "#fff";
                add = false;
                clkArr.current = clkArr.current.filter(e => e.name != elem.name || e.server != elem.server);
            } else {
                target.style.backgroundColor = "#b5b4b1";
                clkArr.current = [...clkArr.current, elem];
            }

            if (type == States.Manage) {
                behavManage(elem.name, elem.server, force, format, add)
            }
        }

    }
    
    function makeEntries(entries) {
 
        return (
            <table id="entries" cellPadding="0" cellSpacing="0" >
                <tbody>
                <tr>
                    <th>Name</th>
                    <th>At</th>
                    <th>Date</th>
                </tr>
                {entries.map(
                    (e, i) => {
                        if (clkArr.current.some(d => e.name == d.name && e.server == d.server)) {
                            return (
                                <tr key={i} onClick={(e) => selectRow(e, force, format)} style={{background:"#b5b4b1"}}>
                                    <td>{e.name}</td>
                                    <td>{e.server}</td>
                                    <td>{e.date}</td>
                                </tr>
                            )
                        } else {
                            return (
                                <tr key={i} onClick={(e) => selectRow(e, force, format)} style={{background:"#fff"}}>
                                    <td>{e.name}</td>
                                    <td>{e.server}</td>
                                    <td>{e.date}</td>
                                </tr>
                            )
                        }
                    }
                )}
                </tbody>
            </table>
        )
    }

    return (
        <>
            <p className="selectText">Select Entry</p>
            <div className="profList">
                {makeEntries(entries)}
            </div>
        </>
    )
}
   
export default ProfList;