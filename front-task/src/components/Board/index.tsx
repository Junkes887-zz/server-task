import React, { useState, useEffect, useRef } from 'react'
import produce from 'immer'

import './styles.css'

import List from '../List'
import api from '../../services/api'
import BoardContext from './context'

interface Task {
    id: number,
    name: string,
    description: string,
    IDStatus: number
}

interface Status {
    id: number,
    name: string,
    Tasks: Task[],
}

const Board = () => {
    const [status, setStatus] = useState<Status[]>([])
    const [list, setList] = useState<Status[]>([])
    const [statusSelected, setStatusSelected] = useState(Number)
    
    useEffect(() =>{
        api.get("status").then(response => {
            setStatus(response.data)
            setList(response.data)
        })
    },[])

    function validations(idNewStatus: number, idOldSatus: number, idTask: number) {
        if(status.filter(i => i.id === idNewStatus)[0] === null) {
            return false
        }
        if(status.filter(i => i.id === idNewStatus)[0].Tasks.filter(i => i.id === idTask).length !== 0) {
            return false
        }
        if(status.filter(i => i.id === idOldSatus)[0] === null) {
            return false
        }
        return true
    }

    function move(idNewStatus: number, idOldSatus: number, task: Task, index: number) {
        if (idNewStatus  === statusSelected) {
            return
        }

        if (validations(idNewStatus, idOldSatus, task.id) === false) {
            return
        }
        status.filter(i => i.id === idOldSatus)[0].Tasks.splice(index, 1)
        status.filter(i => i.id === idNewStatus)[0].Tasks.push(task)
        setStatus([...status])

        task.IDStatus = idNewStatus
        api.put("task", task).then(() => {
            setStatusSelected(idNewStatus)
        })
    }

    function addTask() {
        const data = {
            name: "",
            description: "",
            IDStatus: status[0].id
        }
        api.post<Task>("task", data).then(response => {
            status[0].Tasks.push(response.data)
            setStatus([...status])
        })
    }

    function removeTask(id: number) {
        api.delete("task/"+id).then(() => {
            api.get("status").then(response => {
                setStatus(response.data)
                setList(response.data)
            })
        })
    }

    return (
        <BoardContext.Provider value={{ list, move, addTask, removeTask }}>
            <div className="board">
                { status.map((status, index) => <List key={status.id} data={status} showButton={index === 0}/>) }
            </div>
        </BoardContext.Provider>
    )
}

export default Board