import axios from 'axios';

export const Graphs = {
    Minions : 0,
    Level : 1,
    Xp : 2,
    Gold : 3,
    DmgDealt : 4,
    DmgTaken : 5,
}

// cache data
const dataMap = new Map();
// change host to your own here
const host = 'http://localhost:8080';

async function behavEntries(setEntries) {
    const response = await axios.get(host + '/api/v1/profile');
    setEntries(response.data.data);
};

function behavDelete(selected, force, format) {
    
    let arr = [];
    selected.forEach( e => arr.push({server : e.server, name : e.name}))
    axios.put(host + '/api/v1/profile/delete', {
        list : arr
    }).then(function (response) {
        selected.forEach( e => 
            behavManageRemove(e.name, e.server, force, format)
        )
    }).catch(function (error) {
        console.log(error);
    })
}

function behavUpdate(selected, format) {
    let arr = [];
    selected.forEach( e => arr.push({server : e.server, name : e.name}))
    axios.put(host + '/api/v1/profile', {
        list : arr
    }).catch(function (error) {
        console.log(error);
    })
}

function setGraphs(data, name, format) {
    let minions = setGraph(data.Minions, name, format, Graphs.Minions);
    let level = setGraph(data.Level, name, format, Graphs.Level);
    let xp = setGraph(data.Xp, name, format, Graphs.Xp);
    let gold = setGraph(data.Gold, name, format, Graphs.Gold);
    let dmgDealt = setGraph(data.DmgDealt, name, format, Graphs.DmgDealt);
    let dmgTaken = setGraph(data.DmgTaken, name, format, Graphs.DmgTaken);

    dataMap.set(name, {
        Minions : minions.store, Level : level.store, Xp : xp.store,
        Gold : gold.store, DmgDealt : dmgDealt.store, DmgTaken : dmgTaken.store
    });
    format.current = [minions.add, level.add, xp.add, gold.add, dmgDealt.add, dmgTaken.add];
}

function setGraphStored(name, format, data) {
    let minions = setGraphPartial(data.Minions, name, format, Graphs.Minions);
    let level = setGraphPartial(data.Level, name, format, Graphs.Level);
    let xp = setGraphPartial(data.Xp, name, format, Graphs.Xp);
    let gold = setGraphPartial(data.Gold, name, format, Graphs.Gold);
    let dmgDealt = setGraphPartial(data.DmgDealt, name, format, Graphs.DmgDealt);
    let dmgTaken = setGraphPartial(data.DmgTaken, name, format, Graphs.DmgTaken);

    format.current = [minions, level, xp, gold, dmgDealt, dmgTaken];

}

function setGraphPartial(dataNew, name, format, type) {
    return {
        lines : [...format.current[type].lines, {
            name : name, col : "#" + Math.floor(Math.random()*16777215).toString(16)
        }],
        data : format.current[type].data.concat(dataNew)
    };
}

function setGraph(data, name, format, type) {
    let dataNew = [];
    data.map((i, e) => {
        dataNew.push({
            [name] : i,
            name : e
        })
    });

    return {
        store : dataNew,
        add : {
            lines : [...format.current[type].lines, {
                name : name, col : "#" + Math.floor(Math.random()*16777215).toString(16)
            }],
            data : format.current[type].data.concat(dataNew)
        }
    };
}

function stringToServer(server) {
    switch(server) {
        case "North America":
            return "na1";
        case "Europe":
            return "eun1";
        case "Japan":
            return "jp1";
        case "Korea":
            return "kr1";
        case "Latin America":
            return "la1";
        case "Oceania":
            return "oc1";
        case "Russia":
            return "ru1";
    }
}

function behavAdd(server, name, format, force, added) {

    server = stringToServer(server)
    axios.post(host + '/api/v1/profile', {
            name: name,
            server: server
        })
        .then(function (response) {
            const data = JSON.parse(response.data.data);
            setGraphs(data, name + server, format);
            added.current.push({name : name, server : server});
            console.log(added);
            force();
        }
    ).catch(function (error) {
        console.log(error);
    })
    
}

function behavManage(name, server, force, format, add) {
    if (add) {
        behavManageAdd(name, server, force, format)
    } else {
        behavManageRemove(name, server, force, format)
    }
}

function setGraphRemove(name, format, type) {

    return {
        lines : format.current[type].lines.filter(e => e.name !== name),
        data : format.current[type].data.filter(e => !e.hasOwnProperty(name))
    };
}

function behavManageRemove(name, server, force, format) {
    let mapName = name + server;
    let minions = setGraphRemove(mapName, format, Graphs.Minions);
    let level = setGraphRemove(mapName, format, Graphs.Level);
    let xp = setGraphRemove(mapName, format, Graphs.Xp);
    let gold = setGraphRemove(mapName, format, Graphs.Gold);
    let dmgDealt = setGraphRemove(mapName, format, Graphs.DmgDealt);
    let dmgTaken = setGraphRemove(mapName, format, Graphs.DmgTaken);

    format.current = [minions, level, xp, gold, dmgDealt, dmgTaken];

    force();
}

function behavManageAdd(name, server, force, format) {
    let mapName = name + server;
    if (format.current[Graphs.Minions].lines.find(e => e.name === mapName)) {
        return;
    }
    if (dataMap.has(mapName)) {
        setGraphStored(mapName, format, dataMap.get(mapName));
        force();
    } else {
        axios.get(host + '/api/v1/profile/' + name + "/" + server)
            .then(function (response) {
                const data = JSON.parse(response.data.data);
                setGraphs(data, name + server, format);
                force();
            }
        ).catch(function (error) {
            console.log(error);
        })
    }

}

export {behavUpdate, behavAdd, behavDelete, behavManage, behavEntries};
