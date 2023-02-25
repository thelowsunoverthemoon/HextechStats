import React, { useState } from "react";

import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer
} from "recharts";

import './Charts.css';

function makeLines(format) {
    return (
        <>
            {format.lines.map(
                (l, i) => (
                    <Line key={i} type="monotone" dataKey={l.name} stroke={l.col} />
                )
            )}
        </>
    )
}

function makeChart(format, xLab, yLab, pos) {
    return (
        <div className={pos} >
            {xLab + " vs " + yLab}
            <ResponsiveContainer width="100%" height="40%">
                <LineChart
                    data={format.data}

                >
                <CartesianGrid strokeDasharray="3 3" />
                <XAxis dataKey="name" type="number" domain={['dataMin', 'dataMax']}/>
                <YAxis />
                <Tooltip />
                <Legend />
                {makeLines(format)}
                </LineChart>
            </ResponsiveContainer>
        </div>
    );

}


const Charts = ({format}) => {

    return (

        <div className="chartContainer">
            {makeChart(format.current[3], "Gold", "Min", "topLeft")}
            {makeChart(format.current[0], "Minions", "Min", "topRight")}
            {makeChart(format.current[1], "Level", "Min", "topMid")}
            {makeChart(format.current[2], "Xp", "Min", "botLeft")}
            {makeChart(format.current[4], "Damage Dealt", "Min", "botMid")}
            {makeChart(format.current[5], "Damage Taken", "Min", "botRight")}
        </div>
    )
}

export default Charts;