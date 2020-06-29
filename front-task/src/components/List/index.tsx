import React, { useContext, useRef } from 'react'
import { MdAdd } from 'react-icons/md'
import { useDrop } from 'react-dnd'

import './styles.css'

import Card from '../Card'
import BoardContext from '../Board/context'

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

interface Item {
    type: string,
    task: Task,
    idStatus: number,
    index: number
}

interface IProps {
    data: Status,
    showButton: boolean
}
const List = ({data, showButton}:IProps) => {
    const { move, addTask } = useContext(BoardContext)
    const [, dropRef] = useDrop({
        accept: 'CARD',
        drop(item: Item, monitor) {
            move(data.id, item.idStatus, item.task, item.index)
        }
    })

    return (
        <div ref={dropRef} className="list">
            <header className="header-list">
            <h2>{data.name}</h2>
            { showButton ? (
                <button type="submit" onClick={addTask}>
                    <MdAdd size={24} color="#fff"/>
                </button>
            ) : null}
            </header>

            <ul>
                { data.Tasks.map((task,index) => <Card key={task.id} data={task} index={index} status={data} />) }
            </ul>
        </div>
    )
}

export default List