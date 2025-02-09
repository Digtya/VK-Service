import React, { useEffect, useState } from 'react';
import { Table } from 'antd';
import type { ColumnType } from 'antd/es/table';

interface Container {
    id: number;
    ipAddress: string;
    pingTime: number;
    lastSuccessful: string;
}

const App: React.FC = () => {
    const [containers, setContainers] = useState<Container[]>([]);

    useEffect(() => {
        const fetchContainers = async () => {
            const response = await fetch('http://backend:8080/containers');
            const data = await response.json();
            setContainers(data);
        };

        fetchContainers();
    }, []);

    // Определяем columns как обычный массив
    const columns: ColumnType<Container>[] = [
        {
            title: 'IP Address',
            dataIndex: 'ipAddress',
            key: 'ipAddress',
        },
        {
            title: 'Ping Time (ms)',
            dataIndex: 'pingTime',
            key: 'pingTime',
        },
        {
            title: 'Last Successful',
            dataIndex: 'lastSuccessful',
            key: 'lastSuccessful',
        },
    ];

    return (
        <Table
            dataSource={containers}
            pagination={false}
            columns={columns} // Передаем обычный массив
        />
    );
};

export default App;