import React, { createContext } from 'react'

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

interface IContextProps {
    list: Status[],
    move: (idNewStatus: number, idOldSatus: number, task: Task, index: number) => void,
    addTask: () => void
    removeTask: (id: number) => void
}

export default createContext({} as IContextProps)