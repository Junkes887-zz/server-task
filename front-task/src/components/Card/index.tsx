import React, { ChangeEvent, useState, useEffect, useContext } from 'react'
import { useDrag } from 'react-dnd'
import { MdRemove } from 'react-icons/md'

import './styles.css'

import api from '../../services/api'
import BoardContext from '../Board/context'

interface Task {
    id: number,
    name: string,
    description: string,
    IDStatus: number
}

interface Item {
    type: string,
    task: Task,
    idStatus: number,
    index: number
}

interface Status {
    id: number,
    name: string,
    Tasks: Task[],
}

interface IProps {
    data: Task,
    status: Status,
    index: number,
}

const Card = ({data, status, index}:IProps) => {
    const { removeTask } = useContext(BoardContext)
    const [name, setName] = useState("")
    const [description, setDescription] = useState("")
    const [{isDragging}, dragRef] = useDrag({
        item: { type: 'CARD', task: data, idStatus: status.id, index },
        collect: monitor => ({
            isDragging: monitor.isDragging
        })
    })

    useEffect(() => {
        setName(data.name)
        setDescription(data.description)
    } ,[data])

    function setValueName(e: ChangeEvent<HTMLInputElement>) {
        setName(e.target.value)
    }

    function setValueDescription(e: ChangeEvent<HTMLInputElement>) {
        setDescription(e.target.value)
    }

    const onblur = () => {
        const dto: Task = {
            ...data,
            name,
            description
        }
        api.put("task", dto)
    }

    return (
        <div ref={dragRef} className="card">
            <header>
                <input className="name" value={name} onChange={setValueName} onBlur={onblur}/>
                <button onClick={() => removeTask(data.id)}>
                    <MdRemove size={24} color="#FF0000"/>
                </button>                
            </header>
            <input className="description" value={description} onChange={setValueDescription} onBlur={onblur}/>
        </div>
    )
}

export default Card